package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/lib/errors"
	"github.com/p1ass/midare/lib/logging"
	"go.uber.org/zap"
)

const (
	refreshTokenURL    = "https://api.twitter.com/oauth/request_token"
	authorizationURL   = "https://api.twitter.com/oauth/authenticate"
	accessTokenURL     = "https://api.twitter.com/oauth/access_token"
	twitterAPIEndpoint = "https://api.twitter.com/1.1"
)

var (
	jst = time.FixedZone("Asia/Tokyo", 9*60*60)
)

type client struct {
	consumer    *oauth.Consumer
	callbackURL string
	redisCli    *redis.Client
}

func newClient(consumerKey, consumerSecret, callbackURL string) *client {
	consumer := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   refreshTokenURL,
			AuthorizeTokenUrl: authorizationURL,
			AccessTokenUrl:    accessTokenURL,
		})
	redisCli := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR") + ":6379",
		Password: os.Getenv("REDIS_PASS"),
	})
	return &client{
		consumer:    consumer,
		callbackURL: callbackURL,
		redisCli:    redisCli,
	}
}

// GetRequestTokenAndURL gets a request token associated with the user and login URL
func (cli *client) GetRequestTokenAndURL() (loginURL string, err error) {
	rToken, loginURL, err := cli.consumer.GetRequestTokenAndUrl(cli.callbackURL)
	if err != nil {
		return "", errors.Wrap(err, "failed to get access token")
	}

	v, err := json.Marshal(rToken)
	if err != nil {
		panic(err)
	}
	if err := cli.redisCli.Set(rToken.Token, string(v), 10*time.Minute).Err(); err != nil {
		logging.New().Error("failed to set request token to redis", logging.Error(err))
		return "", err
	}

	return loginURL, nil
}

// AuthorizeToken gets oauth access token
func (cli *client) AuthorizeToken(token, verificationCode string) (*oauth.AccessToken, error) {
	var rToken *oauth.RequestToken
	v, err := cli.redisCli.Get(token).Result()
	if err != nil {
		logging.New().Error("failed to get request token from redis", logging.Error(err))
		return nil, err
	}
	if err := json.Unmarshal([]byte(v), &rToken); err != nil {
		logging.New().Error("failed to unmarshal request token", logging.Error(err))
		return nil, errors.New(errors.Unauthorized, "token secret not found")
	}

	aToken, err := cli.consumer.AuthorizeToken(rToken, verificationCode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to authorize token")
	}

	return aToken, nil
}

// AccountVerifyCredentials fetch twitter profile from twitter api
func (cli *client) AccountVerifyCredentials(token *oauth.AccessToken) (*User, error) {
	httpCli, err := cli.httpClient(token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make http client")
	}

	resp, err := httpCli.Get(twitterAPIEndpoint + "/account/verify_credentials.json")
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch verify credentials from twitter api")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errMsg := &twitterError{}
		err = json.NewDecoder(resp.Body).Decode(errMsg)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode twitter api response")
		}
		return nil, errors.New(errors.Unauthorized, "twitter api response status code=%d message=%v", resp.StatusCode, errMsg.Errors)
	}

	res := &userObject{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode twitter api response")
	}

	twiUser := &User{
		ID:         res.IDStr,
		Name:       res.Name,
		ScreenName: res.ScreenName,
		ImageURL:   res.ProfileImageURL,
	}

	return twiUser, nil
}

func (cli *client) GetUserTweets(token *oauth.AccessToken, screenName, maxID string) ([]*Tweet, error) {
	httpCli, err := cli.httpClient(token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make http client")
	}

	query := url.Values{}
	query.Add("count", "200")
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

	remain, err := strconv.Atoi(resp.Header.Get("X-App-Rate-Limit-Remaining"))
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
		return nil, errors.New(errors.Unauthorized, "twitter api response status code=%d message=%v", resp.StatusCode, errMsg.Errors)
	}

	var res []*tweetObject
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode twitter api response")
	}

	tweets := cli.toTweets(res)
	return tweets, nil
}

// httpClient make *http.Client using access token
func (cli *client) httpClient(aToken *oauth.AccessToken) (*http.Client, error) {
	client, err := cli.consumer.MakeHttpClient(aToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make http client from access token")
	}
	return client, nil
}

type twitterError struct {
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

// userObject is a user object for twitter api
type userObject struct {
	ID              int64  `json:"id"`
	IDStr           string `json:"id_str"`
	Name            string `json:"name"`
	ScreenName      string `json:"screen_name"`
	URL             string `json:"url"`
	ProfileImageURL string `json:"profile_image_url_https"`
}

type tweetObject struct {
	ID         int64  `json:"id"`
	Text       string `json:"text"`
	CreatedStr string `json:"created_at"`
}

func (cli *client) toTweets(tweetObjects []*tweetObject) []*Tweet {
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
