package twitter

import (
	"net/http"
	"time"

	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/datastore"
	"github.com/p1ass/midare/errors"
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

// Client is an interface for calling Twitter api
type Client interface {
	GetLoginURL() (loginURL string, err error)
	AuthorizeToken(token, verificationCode string) (*oauth.AccessToken, error)
	AccountVerifyCredentials(token *oauth.AccessToken) (*TwitterUser, error)
	GetTweets(accessToken *oauth.AccessToken) ([]*Tweet, error)
}

// NewClient returns Client
func NewClient(dsCli datastore.Client) Client {
	twCfg := config.NewTwitter()
	return newClient(twCfg.ConsumerKey, twCfg.ConsumerSecret, twCfg.OAuthCallBackURL, dsCli)
}

type client struct {
	consumer    *oauth.Consumer
	callbackURL string
	dsCli       datastore.Client
}

func newClient(consumerKey, consumerSecret, callbackURL string, dsCli datastore.Client) *client {
	consumer := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   refreshTokenURL,
			AuthorizeTokenUrl: authorizationURL,
			AccessTokenUrl:    accessTokenURL,
		})

	return &client{
		consumer:    consumer,
		callbackURL: callbackURL,
		dsCli:       dsCli,
	}
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
