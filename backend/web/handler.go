package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/p1ass/midare/logging"
	"github.com/p1ass/midare/period"
	"github.com/p1ass/midare/usecase"

	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/twitter"
)

const (
	sessionIDKey = "sessionID"
	sevenDays    = 60 * 60 * 24 * 7
)

// Handler is HTTP handler.
type Handler struct {
	frontendCallbackURL string
	redisCli            *redis.Client
	usecase             *usecase.Usecase
}

// NewHandler returns a new struct of Handler.
func NewHandler(twiCli twitter.Client, frontendCallbackURL string) (*Handler, error) {
	redisCfg := config.ReadRedisConfig()
	redisCli := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr(),
		Password: redisCfg.Password,
	})
	if err := redisCli.Ping().Err(); err != nil {
		logging.New().Error("failed to ping to redis", logging.Error(err))
		return nil, err
	}
	return &Handler{
		frontendCallbackURL: frontendCallbackURL,
		redisCli:            redisCli,
		usecase:             usecase.NewUsecase(twiCli),
	}, nil
}

// GetMe gets my profile.
func (h *Handler) GetMe(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	if accessToken == nil {
		return
	}

	user, err := h.usecase.GetUser(accessToken)
	if err != nil {
		sendError(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAwakePeriods gets awake periods from tweets.
func (h *Handler) GetAwakePeriods(c *gin.Context) {
	type getAwakePeriodsRes struct {
		Periods  []*period.Period `json:"periods"`
		ShareURL string           `json:"shareUrl"`
	}

	accessToken := h.getAccessToken(c)
	if accessToken == nil {
		return
	}

	periods, shareURL, err := h.usecase.GetAwakePeriods(accessToken)
	if err != nil {
		sendError(err, c)
		return
	}

	c.JSON(http.StatusOK, &getAwakePeriodsRes{
		Periods:  periods,
		ShareURL: shareURL,
	})
}
