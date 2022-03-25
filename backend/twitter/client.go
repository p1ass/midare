package twitter

import (
	"fmt"
	"net/http"

	"github.com/g8rswimmer/go-twitter/v2"
	"golang.org/x/oauth2"
)

type client struct {
	cli *twitter.Client
}

func NewClient(token *oauth2.Token) Client {
	c := &twitter.Client{
		Authorizer: newAuthorizer(token),
		Client:     http.DefaultClient,
		Host:       "https://api.twitter.com",
	}
	return &client{
		cli: c,
	}
}

type authorizer struct {
	token *oauth2.Token
}

// newAuthorizer creates OAuth2.0 bearer token authorizer
func newAuthorizer(token *oauth2.Token) *authorizer {
	return &authorizer{token: token}
}

func (a *authorizer) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.token.AccessToken))
}
