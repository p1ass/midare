package twitter

import "github.com/mrjones/oauth"

// Client is a interface for calling twitter api
type Client interface {
	GetRequestTokenAndURL() (loginURL string, err error)
	AuthorizeToken(token, verificationCode string) (*oauth.AccessToken, error)
	AccountVerifyCredentials(token *oauth.AccessToken) (*User, error)
	GetUserTweets(token *oauth.AccessToken, screenName, maxID string) ([]*Tweet, error)
}

// NewClient returns Client
func NewClient(consumerKey, consumerSecret, callbackURL string) Client {
	return newClient(consumerKey, consumerSecret, callbackURL)
}
