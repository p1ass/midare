package web

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/mrjones/oauth"

	"github.com/p1ass/midare/twitter"

	"github.com/gin-gonic/gin"
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
	accessTokens        map[string]*oauth.AccessToken
	mu                  sync.Mutex
}

// NewHandler returns a new struct of Handler.
func NewHandler(twiCli twitter.Client, frontendCallbackURL string, frontendDomain string) *Handler {
	return &Handler{
		twiCli:              twiCli,
		frontendCallbackURL: frontendCallbackURL,
		accessTokens:        map[string]*oauth.AccessToken{},
		mu:                  sync.Mutex{},
	}
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
		// tweets, err := h.twiCli.GetUserTweets(accessToken, "uzimaru0000", maxID)
		tweets, err := h.twiCli.GetUserTweets(accessToken, accessToken.AdditionalData["screen_name"], maxID)
		if err != nil {
			return nil, err
		}
		allTweets = append(allTweets, tweets...)
		if len(allTweets) > 2000 || time.Now().Sub(tweets[len(tweets)-1].Created) > 21*24*time.Hour {
			break
		}
		maxID = allTweets[len(allTweets)-1].ID
	}
	return allTweets, nil
}

func (h *Handler) calcAwakePeriods(ts []*twitter.Tweet) []*period {
	var periods []*period
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
		return nil
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

type period struct {
	OkiTime *twitter.Tweet `json:"okiTime"`
	NeTime  *twitter.Tweet `json:"neTime"`
}
