package web

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/entity"
	"github.com/p1ass/midare/lib/logging"
	"github.com/p1ass/midare/twitter"
	"github.com/p1ass/midare/usecase"
	"github.com/patrickmn/go-cache"
)

const (
	sessionIDKey = "sessionID"
	sevenDays    = 60 * 60 * 24 * 7

	oldestTweetTime = 21 * 24 * time.Hour
)

// Handler ia HTTP handler.
type Handler struct {
	twiCli              twitter.Client
	frontendCallbackURL string
	redisCli            *redis.Client
	responseCache       *cache.Cache
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
		twiCli:              twiCli,
		frontendCallbackURL: frontendCallbackURL,
		redisCli:            redisCli,
		responseCache:       cache.New(5*time.Minute, 5*time.Minute),
		usecase:             usecase.NewUsecase(twiCli),
	}, nil
}

// GetMe gets my profile.
func (h *Handler) GetMe(c *gin.Context) {
	accessToken := h.getAccessToken(c)
	if accessToken == nil {
		return
	}

	user, err := h.twiCli.AccountVerifyCredentials(accessToken)
	if err != nil {
		sendError(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAwakePeriods gets awake periods from tweets.
func (h *Handler) GetAwakePeriods(c *gin.Context) {
	type getAwakePeriodsRes struct {
		Periods  []*entity.Period `json:"periods"`
		ShareURL string           `json:"shareUrl"`
	}

	accessToken := h.getAccessToken(c)
	if accessToken == nil {
		return
	}

	screenName := accessToken.AdditionalData["screen_name"]

	cached, ok := h.responseCache.Get(screenName)
	if ok {
		c.JSON(http.StatusOK, cached.(*getAwakePeriodsRes))
		return
	}

	tweets, err := h.usecase.GetTweets(accessToken)
	if err != nil {
		sendError(err, c)
		return
	}

	periods := h.usecase.CalcAwakePeriods(tweets)

	shareID := uuid.New().String()

	url := h.usecase.UploadImage(periods, shareID, accessToken)

	res := &getAwakePeriodsRes{Periods: periods, ShareURL: url}

	h.responseCache.SetDefault(screenName, res)

	c.JSON(http.StatusOK, res)
}
