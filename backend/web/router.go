package web

import (
	"net/http"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/logging"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// NewRouter returns a gin router
func NewRouter(handler *Handler, allowOrigin string) (*gin.Engine, error) {
	r := gin.New()

	r.Use(gin.Recovery())

	logger := logging.New()
	r.Use(ginzap.RecoveryWithZap(logger, true))

	store := cookie.NewStore([]byte(config.ReadSessionKey()))
	store.Options(sessions.Options{
		MaxAge:   86400 * 7,
		Secure:   !config.IsLocal(),
		HttpOnly: true,
	})
	r.Use(sessions.Sessions("session-store", store))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowOrigin},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Cookie", "Content-Type", "Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/login", handler.StartSignInWithTwitter)
	r.GET("/callback", handler.TwitterCallback)

	withAuthGrp := r.Group("/")
	withAuthGrp.Use(AuthMiddleware())
	withAuthGrp.GET("/me", handler.GetMe)
	withAuthGrp.GET("/periods", handler.GetAwakePeriods)

	return r, nil
}
