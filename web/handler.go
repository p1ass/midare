package web

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/lib/errors"
	"github.com/p1ass/midare/lib/logging"
	"github.com/p1ass/midare/twitter"
)

const (
	sessionIDKey = "sessionID"
	sevenDays    = 60 * 60 * 24 * 7

	// この時間以内にツイートされていたらその時間は起きていることにする
	awakeThreshold = 3*time.Hour + 30*time.Minute
)

// Handler ia HTTP handler.
type Handler struct {
	twiCli              twitter.Client
	frontendCallbackURL string
	mu                  sync.Mutex
	redisCli            *redis.Client
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
		mu:                  sync.Mutex{},
		redisCli:            redisCli,
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
	accessToken := h.getAccessToken(c)
	if accessToken == nil {
		return
	}

	tweets, err := h.getTweets(accessToken)
	if err != nil {
		sendError(err, c)
		return
	}

	type getAwakePeriodsRes struct {
		Periods []*period `json:"periods"`
	}

	c.JSON(http.StatusOK, &getAwakePeriodsRes{Periods: h.calcAwakePeriods(tweets)})
}

// getTweets gets more than 1000 tweets.
func (h *Handler) getTweets(accessToken *oauth.AccessToken) ([]*twitter.Tweet, error) {
	var allTweets []*twitter.Tweet
	maxID := ""
	for {
		tweets, err := h.twiCli.GetUserTweets(accessToken, accessToken.AdditionalData["screen_name"], maxID)
		if err != nil {
			return nil, err
		}
		if len(tweets) == 0 {
			return []*twitter.Tweet{}, nil
		}
		filtered := h.filterByCreated(tweets)
		allTweets = append(allTweets, filtered...)
		if len(allTweets) > 1500 || len(filtered) < len(tweets) {
			break
		}
		maxID = allTweets[len(allTweets)-1].ID
	}
	return allTweets, nil
}

func (h *Handler) filterByCreated(tweets []*twitter.Tweet) []*twitter.Tweet {
	var filtered []*twitter.Tweet

	for _, t := range tweets {
		if time.Now().Sub(t.Created) <= 21*24*time.Hour {
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
	if neTweet == nil {
		return periods
	}

	for _, t := range ts[startIdx:] {
		if h.containExcludeWord(t.Text) {
			continue
		}

		durationBetweenTweets := lastTweet.Created.Sub(t.Created)
		if durationBetweenTweets < awakeThreshold {
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

	return periods
}

func (h *Handler) containExcludeWord(text string) bool {
	excludeWords := []string{"ぼくへ 生活習慣乱れてませんか？", "#contributter_report", "のポスト数"}
	for _, word := range excludeWords {
		if strings.Contains(text, word) {
			return true
		}
	}
	return false
}

func (h *Handler) UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		sendError(err, c)
		return
	}
	if header.Size > 1000000 {
		sendServiceError(&errors.ServiceError{
			Code: http.StatusBadRequest,
		}, c)
		return
	}
	binary, err := ioutil.ReadAll(file)
	if err != nil {
		sendError(err, c)
		return
	}

	fileName := uuid.New().String()
	fileNameWithExt := fileName + ".jpeg"

	if err := h.uploadImageToCloudStorage(fileNameWithExt, binary); err != nil {
		sendError(err, c)
		return
	}

	type response struct {
		ShareURL string `json:"shareUrl"`
	}
	res := &response{
		ShareURL: os.Getenv("CORS_ALLOW_ORIGIN") + "/share/" + fileName,
	}
	c.JSON(http.StatusOK, res)
	return
}

func (h *Handler) uploadImageToCloudStorage(fileName string, image []byte) error {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return errors.Wrap(err, "new storage client")
	}

	bucketName := os.Getenv("BUCKET_NAME")
	bucket := client.Bucket(bucketName)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	wc := bucket.Object(fileName).NewWriter(ctx)
	if _, err = io.WriteString(wc, string(image)); err != nil {
		return errors.Wrap(err, "upload image: %s", fileName)
	}
	if err := wc.Close(); err != nil {
		return errors.Wrap(err, "close writer: %s", fileName)
	}

	return nil
}

type period struct {
	OkiTime *twitter.Tweet `json:"okiTime"`
	NeTime  *twitter.Tweet `json:"neTime"`
}
