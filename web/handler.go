package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/mrjones/oauth"

	"github.com/p1ass/midare/lib/errors"
	"github.com/p1ass/midare/lib/logging"
	"github.com/p1ass/midare/twitter"
	"github.com/patrickmn/go-cache"
)

const (
	sessionIDKey = "sessionID"
	sevenDays    = 60 * 60 * 24 * 7

	// この時間以内にツイートされていたらその時間は起きていることにする
	awakeThreshold = 3*time.Hour + 30*time.Minute

	oldestTweetTime = 21 * 24 * time.Hour
)

// Handler ia HTTP handler.
type Handler struct {
	twiCli              twitter.Client
	frontendCallbackURL string
	redisCli            *redis.Client
	responseCache       *cache.Cache
}

// NewHandler returns a new struct of Handler.
func NewHandler(twiCli twitter.Client, frontendCallbackURL string) (*Handler, error) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR") + ":6379",
		Password: os.Getenv("REDIS_PASS"),
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
		Periods  []*period `json:"periods"`
		ShareURL string    `json:"shareUrl"`
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

	tweets, err := h.getTweets(accessToken)
	if err != nil {
		sendError(err, c)
		return
	}

	periods := h.calcAwakePeriods(tweets)

	shareID := uuid.New().String()

	url := h.uploadImage(periods, shareID)

	res := &getAwakePeriodsRes{Periods: periods, ShareURL: url}

	h.responseCache.SetDefault(screenName, res)

	c.JSON(http.StatusOK, res)
}

// getTweets gets more than 1000 tweets.
func (h *Handler) getTweets(accessToken *oauth.AccessToken) ([]*twitter.Tweet, error) {
	screenName := accessToken.AdditionalData["screen_name"]

	var allTweets []*twitter.Tweet
	maxID := ""
	for {
		tweets, err := h.twiCli.GetUserTweets(accessToken, screenName, maxID)
		if err != nil {
			return nil, err
		}
		if len(tweets) == 0 {
			return []*twitter.Tweet{}, nil
		}
		filtered := h.filterByCreated(tweets)
		allTweets = append(allTweets, filtered...)
		if len(allTweets) > 2000 || len(filtered) < len(tweets) {
			break
		}
		maxID = allTweets[len(allTweets)-1].ID
	}

	return allTweets, nil
}

func (h *Handler) filterByCreated(tweets []*twitter.Tweet) []*twitter.Tweet {
	var filtered []*twitter.Tweet

	for _, t := range tweets {
		if time.Now().Sub(t.Created) <= oldestTweetTime {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

func (h *Handler) calcAwakePeriods(ts []*twitter.Tweet) []*period {
	periods := []*period{}
	var neTweet *twitter.Tweet
	var okiTweet *twitter.Tweet
	var lastTweet *twitter.Tweet
	startIdx := 1
	for i, t := range ts {
		if !h.containExcludeWord(t.Text) {
			neTweet = t
			okiTweet = t
			lastTweet = t
			startIdx = i + 1
			break
		}
	}
	if lastTweet == nil {
		return periods
	}

	for _, t := range ts[startIdx:] {
		if h.containExcludeWord(t.Text) {
			continue
		}

		durationBetweenTweets := lastTweet.Created.Sub(t.Created)
		if durationBetweenTweets <= awakeThreshold {
			okiTweet = t
			lastTweet = t
			continue
		}

		if okiTweet != neTweet {
			periods = append(periods, &period{
				OkiTime: okiTweet,
				NeTime:  neTweet,
			})
		}

		okiTweet = t
		neTweet = t
		lastTweet = t
	}

	if okiTweet != neTweet {
		periods = append(periods, &period{
			OkiTime: okiTweet,
			NeTime:  neTweet,
		})
	}

	return periods
}

func (h *Handler) containExcludeWord(text string) bool {
	excludeWords := []string{"ぼくへ 生活習慣乱れてませんか？", "みんなへ 生活習慣乱れてませんか？", "#contributter_report", "のポスト数"}
	for _, word := range excludeWords {
		if strings.Contains(text, word) {
			return true
		}
	}
	return false
}

func (h *Handler) uploadImage(periods []*period, shareID string) string {

	go h.uploadImageThroughCloudFunctions(shareID, periods)

	return os.Getenv("CORS_ALLOW_ORIGIN") + "/share/" + shareID
}

func (h *Handler) uploadImageThroughCloudFunctions(uuid string, periods []*period) error {
	type request struct {
		UUID    string    `json:"uuid"`
		Periods []*period `json:"periods"`
	}

	req := &request{
		UUID:    uuid,
		Periods: periods,
	}
	encoded, _ := json.Marshal(req)

	_, err := http.Post(os.Getenv("CLOUD_FUNCTIONS_URL"), "application/json", bytes.NewBuffer(encoded))
	if err != nil {
		logging.New().Error(err.Error())
		return errors.Wrap(err, "post period data to cloud functions")
	}

	return nil
}

type period struct {
	OkiTime *twitter.Tweet `json:"okiTime"`
	NeTime  *twitter.Tweet `json:"neTime"`
}
