package web

import (
	"net/http"

	"github.com/p1ass/midare/errors"
	"github.com/p1ass/midare/logging"

	"github.com/mrjones/oauth"

	"github.com/gin-gonic/gin"
)

// StartSignInWithTwitter start twitter oauth sign in
func (h *Handler) StartSignInWithTwitter(c *gin.Context) {
	url, err := h.usecase.GetLoginUrl()
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
		logger.Error("oauth token should be not empty")
		c.Redirect(http.StatusFound, h.frontendCallbackURL)
		return
	}

	ov := c.DefaultQuery("oauth_verifier", "")
	if ov == "" {
		logger.Error("oauth verifier should be not empty")
		c.Redirect(http.StatusFound, h.frontendCallbackURL)
		return
	}

	accessToken, err := h.usecase.AuthorizeToken(token, ov)
	if err != nil {
		logger.Error("failed to authorize", logging.Error(err))
		c.Redirect(http.StatusFound, h.frontendCallbackURL)
		return
	}

	twiUser, err := h.usecase.GetUser(accessToken)
	if err != nil {
		logger.Error("failed to get twitter user", logging.Error(err))
		c.Redirect(http.StatusFound, h.frontendCallbackURL)
		return
	}

	if err := h.dsCli.StoreAccessToken(c.Request.Context(), twiUser.ID, accessToken); err != nil {
		logger.Error("failed to save access token", logging.Error(err))
		c.Redirect(http.StatusFound, h.frontendCallbackURL)
	}

	if err := setSessionAndCookie(c, twiUser.ID); err != nil {
		sendError(errors.Wrap(err, "failed to set session"), c)
		return
	}
	c.Redirect(http.StatusFound, h.frontendCallbackURL)
}

func (h *Handler) getAccessToken(c *gin.Context) *oauth.AccessToken {
	v, ok := c.Get(userIDContextKey)
	if !ok {
		sendServiceError(errors.NewUnknown("user id must be set with context"), c)
		return nil
	}
	userID := v.(string)

	logger := logging.New()

	accessToken, err := h.dsCli.FetchAccessToken(c.Request.Context(), userID)
	if err != nil {
		logger.Info("failed to get access token", logging.Error(errors.Cause(err)))
		sendServiceError(&errors.ServiceError{Code: errors.Unauthorized}, c)
		return nil
	}

	return accessToken
}
