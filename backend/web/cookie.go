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

func setSessionAndCookie(c *gin.Context, userID string) error {
	session := sessions.Default(c)
	sessID := crypto.LongSecureRandomBase64()
	session.Set(sessID, userID)
	err := session.Save()
	if err != nil {
		return errors.Wrap(err, "failed to save session")
	}

	c.SetCookie(sessionIDKey, sessID, sevenDays, "/", "", !config.IsLocal(), true)
	return nil
}

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

func getOAuthStateID(c *gin.Context) (string, error) {
	session := sessions.Default(c)
	stateID, ok := session.Get(oauthStateKey).(string)
	if ok {
		return stateID, nil
	}
	stateID = crypto.LongSecureRandomBase64()
	session.Set(oauthStateKey, stateID)
	err := session.Save()
	if err != nil {
		return "", errors.Wrap(err, "failed to save oauth state key")
	}

	return stateID, nil
}
