package web

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/crypto"
	"github.com/p1ass/midare/errors"
)

const (
	sessionIDKey  = "sessionID"
	oauthStateKey = "oauthStateID"
)

// setSessionAndCookie creates login session and saves it to cookie.
func setSessionAndCookie(c *gin.Context, userID string) error {
	session := sessions.Default(c)
	sessID := crypto.SecureRandomBase64Encoded(64)
	session.Set(sessID, userID)
	err := session.Save()
	if err != nil {
		return errors.Wrap(err, "failed to save session")
	}

	c.SetCookie(sessionIDKey, sessID, sevenDays, "/", "", !config.IsLocal(), true)
	return nil
}

// getUserIDFromCookie gets logged in userID from login session.
func getUserIDFromCookie(c *gin.Context) (string, error) {
	sessID, err := c.Cookie(sessionIDKey)
	if err != nil {
		return "", errors.New(errors.Unauthorized, "must be logged in")
	}

	session := sessions.Default(c)
	userID, ok := session.Get(sessID).(string)
	if !ok {
		return "", errors.New(errors.Unauthorized, http.StatusText(http.StatusUnauthorized))
	}
	return userID, nil
}

// getOAuthStateID returns id associated with User Agent.
// It is used for identifying OAuth2 state.
// State is used before completing OAuth2 flow, so It is independent to login session.
func getOAuthStateID(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	stateID, ok := session.Get(oauthStateKey).(string)
	if ok {
		return stateID, nil
	}
	stateID = crypto.SecureRandomBase64Encoded(64)
	session.Set(oauthStateKey, stateID)
	err := session.Save()
	if err != nil {
		return "", errors.Wrap(err, "failed to save oauth state key")
	}

	return stateID, nil
}
