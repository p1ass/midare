package web

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/p1ass/midare/lib/crypto"
	"github.com/p1ass/midare/lib/errors"
)

func setSessionAndCookie(c *gin.Context, userID string, frontendDomain string) error {
	session := sessions.Default(c)
	sessID := crypto.LongSecureRandomBase64()
	session.Set(sessID, userID)
	err := session.Save()
	if err != nil {
		return errors.Wrap(err, "failed to save session")
	}

	c.SetCookie(sessionIDKey, sessID, sevenDays, "/", frontendDomain, false, true)
	c.SetSameSite(http.SameSiteNoneMode)
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
