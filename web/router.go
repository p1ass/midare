package web

import (
	"net/http"
	"os"

	"github.com/p1ass/midare/lib/logging"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// NewRouter returns a gin router
func NewRouter(twiHandler *Handler, allowOrigin string) (*gin.Engine, error) {
	r := gin.New()

	r.Use(gin.Recovery())

	logger := logging.New()
	// r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	store, err := redis.NewStore(256, "tcp", os.Getenv("REDIS_ADDR")+":6379", os.Getenv("REDIS_PASS"), []byte("secret"))
	if err != nil {
		logging.New().Error("failed to prepare redis", logging.Error(err))
		return nil, err
	}
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

	r.GET("/login", twiHandler.StartSignInWithTwitter)
	r.GET("/callback", twiHandler.TwitterCallback)

	withAuthGrp := r.Group("/")
	withAuthGrp.Use(AuthMiddleware())
	withAuthGrp.GET("/me", twiHandler.GetMe)
	withAuthGrp.GET("/periods", twiHandler.GetAwakePeriods)
	withAuthGrp.POST("images", twiHandler.UploadImage)

	return r, nil
}
