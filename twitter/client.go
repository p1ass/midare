package twitter

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/mrjones/oauth"
	"github.com/p1ass/seikatsu-syukan-midare/lib/errors"
)

const (
	refreshTokenURL    = "https://api.twitter.com/oauth/request_token"
	authorizationURL   = "https://api.twitter.com/oauth/authenticate"
	accessTokenURL     = "https://api.twitter.com/oauth/access_token"
	twitterAPIEndpoint = "https://api.twitter.com/1.1"
)

type client struct {
	consumer      *oauth.Consumer
	callbackURL   string
	requestTokens map[string]*oauth.RequestToken
	mu            sync.Mutex
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
	return &client{
		consumer:      consumer,
		callbackURL:   callbackURL,
		requestTokens: map[string]*oauth.RequestToken{},
		mu:            sync.Mutex{},
	}
}

// GetRequestTokenAndURL gets a request token associated with the user and login URL
func (cli *client) GetRequestTokenAndURL() (loginURL string, err error) {
	rToken, loginURL, err := cli.consumer.GetRequestTokenAndUrl(cli.callbackURL)
	if err != nil {
		return "", errors.Wrap(err, "failed to get access token")
	}

	cli.mu.Lock()
	cli.mu.Unlock()
	cli.requestTokens[rToken.Token] = rToken

	return loginURL, nil
}

// AuthorizeToken gets oauth access token
func (cli *client) AuthorizeToken(token, verificationCode string) (*oauth.AccessToken, error) {
	cli.mu.Lock()
	defer cli.mu.Unlock()
	rToken, ok := cli.requestTokens[token]
	if !ok {
		return nil, errors.New(errors.Unauthorized, "token secret not found")
	}

	aToken, err := cli.consumer.AuthorizeToken(rToken, verificationCode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to authorize token")
	}

	return aToken, nil
}

// AccountVerifyCredentials fetch twitter profile from twitter api
func (cli *client) AccountVerifyCredentials(token, secret string) (*User, error) {
	httpCli, err := cli.httpClient(&oauth.AccessToken{
		Token:  token,
		Secret: secret,
	})
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
