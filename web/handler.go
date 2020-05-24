package web

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/mrjones/oauth"

	"github.com/p1ass/seikatsu-syukan-midare/twitter"

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
	twiCli         twitter.Client
	frontendURL    string
	frontendDomain string
	accessTokens   map[string]*oauth.AccessToken
	mu             sync.Mutex
}

// NewHandler returns a new struct of Handler.
func NewHandler(twiCli twitter.Client, frontendURL string, frontendDomain string) *Handler {
	return &Handler{
		twiCli:         twiCli,
		frontendURL:    frontendURL,
		frontendDomain: frontendDomain,
		accessTokens:   map[string]*oauth.AccessToken{},
		mu:             sync.Mutex{},
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
		tweets, err := h.twiCli.UserTimeLine(accessToken, accessToken.AdditionalData["screen_name"], maxID)
		if err != nil {
			return nil, err
		}
		allTweets = append(allTweets, tweets...)
		if len(allTweets) > 1000 {
			break
		}
		maxID = allTweets[len(allTweets)-1].ID
	}
	return allTweets, nil
}

func (h *Handler) calcAwakePeriods(ts []*twitter.Tweet) []*period {
	var periods []*period
	neTime := ts[0]
	okiTime := ts[0]
	lastTweetTime := ts[0]

	for _, t := range ts[1:] {
		if h.containExcludeWord(t.Text) {
			continue
		}

		durationBetweenTweets := lastTweetTime.Created.Sub(t.Created)
		if durationBetweenTweets < awakeThreshold {
			okiTime = t
			lastTweetTime = t
			continue
		}

		if okiTime != neTime {
			periods = append(periods, &period{
				OkiTime: okiTime,
				NeTime:  neTime,
			})
		}

		neTime = t
		lastTweetTime = t
	}

	return periods
}

func (h *Handler) containExcludeWord(text string) bool {
	excludeWords := []string{"ぼくへ 生活習慣乱れてませんか？", "#contributter_report"}
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
