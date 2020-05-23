package handler

import (
	"net/http"
	"sync"

	"github.com/mrjones/oauth"

	"github.com/p1ass/seikatsu-syukan-midare/lib/errors"
	"github.com/p1ass/seikatsu-syukan-midare/lib/logging"
	"github.com/p1ass/seikatsu-syukan-midare/twitter"

	"github.com/gin-gonic/gin"
)

const (
	sessionIDKey = "sessionID"
	sevenDays    = 60 * 60 * 24 * 7
)

type Handler struct {
	twiCli         twitter.Client
	frontendURL    string
	frontendDomain string
	accessTokens   map[string]*oauth.AccessToken
	mu             sync.Mutex
}

func NewHandler(twiCli twitter.Client, frontendURL string, frontendDomain string) *Handler {
	return &Handler{
		twiCli:         twiCli,
		frontendURL:    frontendURL,
		frontendDomain: frontendDomain,
		accessTokens:   map[string]*oauth.AccessToken{},
		mu:             sync.Mutex{},
	}
}

// StartSignInWithTwitter start twitter oauth sign in
func (h *Handler) StartSignInWithTwitter(c *gin.Context) {
	url, err := h.twiCli.GetRequestTokenAndURL()
	if err != nil {
		sendError(errors.Wrap(err, "failed to get redirect url"), c)
		return
	}

	c.Header("Cache-Control", "no-cache")
	c.Header("Pragma", "no-cache")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// TwitterCallback handles callback function after login succeeded
// Redirect to frontend even if callback function fails
func (h *Handler) TwitterCallback(c *gin.Context) {
	logger := logging.New()

	token := c.DefaultQuery("oauth_token", "")
	if token == "" {
		logger.Warn("oauth token should be not empty")
		c.Redirect(http.StatusFound, h.frontendURL+"/callback")
		return
	}

	ov := c.DefaultQuery("oauth_verifier", "")
	if ov == "" {
		logger.Warn("oauth verifier should be not empty")
		c.Redirect(http.StatusFound, h.frontendURL+"/callback")
		return
	}

	accessToken, err := h.twiCli.AuthorizeToken(token, ov)
	if err != nil {
		logger.Warn("failed to authorize", logging.Error(err))
		c.Redirect(http.StatusFound, h.frontendURL+"/callback")
		return
	}

	twiUser, err := h.twiCli.AccountVerifyCredentials(accessToken)
	if err != nil {
		logger.Warn("failed to get twitter user", logging.Error(err))
		c.Redirect(http.StatusFound, h.frontendURL+"/callback")
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()
	h.accessTokens[twiUser.ID] = accessToken
	if err := setSessionAndCookie(c, twiUser.ID, h.frontendDomain); err != nil {
		sendError(errors.Wrap(err, "failed to set session"), c)
		return
	}
	c.Redirect(http.StatusFound, h.frontendURL+"/callback")
}

func (h *Handler) GetMe(c *gin.Context) {
	v, ok := c.Get(userIDContextKey)
	if !ok {
		sendServiceError(errors.NewUnknown("user id must be set with context"), c)
		return
	}
	userID := v.(string)

	accessToken, ok := h.accessTokens[userID]
	if !ok {
		sendServiceError(&errors.ServiceError{Code: errors.Unauthorized}, c)
		return
	}

	user, err := h.twiCli.AccountVerifyCredentials(accessToken)
	if err != nil {
		sendError(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetTweets(c *gin.Context) {
	v, ok := c.Get(userIDContextKey)
	if !ok {
		sendServiceError(errors.NewUnknown("user id must be set with context"), c)
		return
	}
	userID := v.(string)

	accessToken, ok := h.accessTokens[userID]
	if !ok {
		sendServiceError(&errors.ServiceError{Code: errors.Unauthorized}, c)
		return
	}

	tweets, err := h.twiCli.UserTimeLine(accessToken, accessToken.AdditionalData["user_id"])
	if err != nil {
		sendError(err, c)
		return
	}
	c.JSON(http.StatusOK, tweets)
}
