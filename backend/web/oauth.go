package web

import (
	"net/http"

	"github.com/p1ass/midare/errors"
	"github.com/p1ass/midare/logging"
	"golang.org/x/oauth2"

	"github.com/gin-gonic/gin"
)

// StartSignInWithTwitter start twitter OAuth2 authorization code flow
func (h *Handler) StartSignInWithTwitter(c *gin.Context) {
	stateID, err := getOAuthStateID(c)
	if err != nil {
		sendError(errors.Wrap(err, "failed to get oauth state id"), c)
		return
	}
	url, err := h.usecase.GetLoginUrl(stateID)
	if err != nil {
		sendError(errors.Wrap(err, "failed to get redirect url"), c)
		return
	}

	c.Header("Cache-Control", "no-cache")
	c.Header("Pragma", "no-cache")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// TwitterCallback handles callback function after OAuth2 use authorization
// Redirect to frontend even if callback function fails
func (h *Handler) TwitterCallback(c *gin.Context) {
	logger := logging.New()

	code := c.DefaultQuery("code", "")
	if code == "" {
		logger.Error("code should be not empty")
		c.Redirect(http.StatusFound, h.frontendCallbackURL)
		return
	}

	state := c.DefaultQuery("state", "")
	if state == "" {
		logger.Error("state should be not empty")
		c.Redirect(http.StatusFound, h.frontendCallbackURL)
		return
	}

	stateID, err := getOAuthStateID(c)
	if err != nil {
		sendError(errors.Wrap(err, "failed to get oauth state id"), c)
		return
	}
	user, err := h.usecase.AuthorizeToken(c.Request.Context(), stateID, code, state)
	if err != nil {
		logger.Error("failed to authorize", logging.Error(err))
		c.Redirect(http.StatusFound, h.frontendCallbackURL)
		return
	}

	if err := setSessionAndCookie(c, user.ID); err != nil {
		sendError(errors.Wrap(err, "failed to set session"), c)
		return
	}
	c.Redirect(http.StatusFound, h.frontendCallbackURL)
}

// getAccessToken gets OAuth2 access token from datastore.
// It is expected that context passed AuthMiddleware
// TODO: usecaseに移しても良いかも
func (h *Handler) getAccessToken(c *gin.Context) (string, *oauth2.Token) {
	v, ok := c.Get(userIDContextKey)
	if !ok {
		sendServiceError(errors.NewUnknown("user id must be set with context"), c)
		return "", nil
	}
	userID := v.(string)

	logger := logging.New()

	accessToken, err := h.dsCli.FetchAccessToken(c.Request.Context(), userID)
	if err != nil {
		if se, ok := errors.Cause(err).(*errors.ServiceError); ok && se.Code == errors.NotFound {
			logger.Info("access token not found")
		} else {
			logger.Error("failed to get access token", logging.Error(errors.Cause(err)))
		}
		sendServiceError(&errors.ServiceError{Code: errors.Unauthorized}, c)
		return "", nil
	}

	return userID, accessToken
}
