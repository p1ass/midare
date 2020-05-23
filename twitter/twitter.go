package twitter

import (
	"sync"

	"github.com/mrjones/oauth"
)

// Client is a interface for using twitter api
type Client interface {
	GetRequestTokenAndURL() (loginURL string, err error)
	AuthorizeToken(token, verificationCode string) (*Credential, error)
	AccountVerifyCredentials(token, secret string) (*User, error)
	FriendsList(cred Credential) ([]*UserObject, error)
	GetUserByScreenName(screenName string, cred Credential) (*UserObject, error)
}

const (
	refreshTokenURL         = "https://api.twitter.com/oauth/request_token"
	authorizationURL        = "https://api.twitter.com/oauth/authenticate"
	accessTokenURL          = "https://api.twitter.com/oauth/access_token"
	twitterAPIEndpoint      = "https://api.twitter.com/1.1"
	requestTokenRedisPrefix = "request-token-"
)

// UserObject is a user object for twitter api
type UserObject struct {
	ID              int64  `json:"id"`
	IDStr           string `json:"id_str"`
	Name            string `json:"name"`
	ScreenName      string `json:"screen_name"`
	URL             string `json:"url"`
	ProfileImageURL string `json:"profile_image_url_https"`
}

// NewTwitterClient returns new struct of Client
func NewTwitterClient(consumerKey, consumerSecret, callbackURL string) (Client, error) {
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
	}, nil
}
