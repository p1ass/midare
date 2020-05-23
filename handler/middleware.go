package handler

import "github.com/gin-gonic/gin"

const userIDContextKey = "userID"

// AuthMiddleware get session id from cookie and set user id to context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := getUserIDFromCookie(c)
		if err != nil {
			sendError(err, c)
			c.Abort()
			return
		}
		c.Set(userIDContextKey, userID)
		c.Next()
	}
}
