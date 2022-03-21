package twitterv1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"

	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/errors"
	"github.com/p1ass/midare/logging"
	"go.uber.org/zap"
)

const (
	maxElapsedDuration = 21 * 24 * time.Hour

	tweetsCountPerAPI = 200
)

// GetTweets gets 2000 tweets.
func (cli *client) GetTweets(accessToken *oauth.AccessToken) ([]*Tweet, error) {
	screenName := accessToken.AdditionalData["screen_name"]

	var mergedTweets []*Tweet
	maxID := ""
	// 一度のAPIで200件取得するので最大2000件になる
	for i := 0; i < 2000/tweetsCountPerAPI; i++ {
		tweets, err := cli.getUserTweets(accessToken, screenName, maxID)
		if err != nil {
			return nil, err
		}
		if len(tweets) == 0 {
			return []*Tweet{}, nil
		}
		extracted := extractWithinMaxElapsedDuration(tweets)
		mergedTweets = append(mergedTweets, extracted...)
		if doesReachFirstTweet(tweets) || exceededMaxElapsed(extracted, tweets) {
			break
		}
		maxID = mergedTweets[len(mergedTweets)-1].ID
	}

	return mergedTweets, nil
}

func (cli *client) getUserTweets(token *oauth.AccessToken, screenName, maxID string) ([]*Tweet, error) {
	httpCli, err := cli.httpClient(token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make http client")
	}

	query := url.Values{}
	query.Add("count", strconv.Itoa(tweetsCountPerAPI))
	query.Add("screen_name", screenName)
	query.Add("exclude_replies", "false")
	query.Add("trim_user", "true")
	if maxID != "" {
		query.Add("max_id", maxID)
	}

	resp, err := httpCli.Get(twitterAPIEndpoint + "/statuses/user_timeline.json?" + query.Encode())
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch verify credentials from twitter api")
	}
	defer resp.Body.Close()

	remain, _ := strconv.Atoi(resp.Header.Get("X-App-Rate-Limit-Remaining"))
	if remain%100 == 0 || remain <= 10000 {
		logging.New().Info(resp.Status+resp.Header.Get("X-App-Rate-Limit-Limit")+":"+resp.Header.Get("X-App-Rate-Limit-Remaining")+":"+resp.Header.Get("X-App-Rate-Limit-Reset"), zap.Int("rate_limit", remain))
	}

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		logging.New().Error(resp.Status + resp.Header.Get(" x-rate-limit-limit,") + ":" + resp.Header.Get("x-rate-limit-remaining ") + ":" + resp.Header.Get("x-rate-limit-reset") + string(body))
		errMsg := &twitterError{}
		err = json.Unmarshal(body, &errMsg)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode twitter api response")
		}
		return nil, errors.New(errors.Unknown, "twitter api response status code=%d message=%v", resp.StatusCode, errMsg.Errors)
	}

	var res []*tweetObject
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode twitter api response")
	}

	tweets := toTweets(res)
	return tweets, nil
}

func toTweets(tweetObjects []*tweetObject) []*Tweet {
	var ts []*Tweet

	for _, t := range tweetObjects {
		created, _ := time.Parse(time.RubyDate, t.CreatedStr)
		ts = append(ts, &Tweet{
			ID:      fmt.Sprintf("%d", t.ID),
			Text:    t.Text,
			Created: created.In(jst),
		})
	}
	return ts
}

func exceededMaxElapsed(extracted, tweets []*Tweet) bool {
	return len(extracted) < len(tweets)
}

func doesReachFirstTweet(tweets []*Tweet) bool {
	return len(tweets) <= 1
}

func extractWithinMaxElapsedDuration(tweets []*Tweet) []*Tweet {
	var filtered []*Tweet

	for _, t := range tweets {
		if time.Since(t.Created) <= maxElapsedDuration {
			filtered = append(filtered, t)
		}
	}
	return filtered
}

type tweetObject struct {
	ID         int64  `json:"id"`
	Text       string `json:"text"`
	CreatedStr string `json:"created_at"`
}
