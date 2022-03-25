package twitter

import (
	"context"

	"golang.org/x/oauth2"
)

// Client is a Twitter client. It must be created per bearer token.
type Client interface {
	GetMe(ctx context.Context) (*User, error)
	GetTweets(ctx context.Context, userID string) ([]*Tweet, error)
}

// Auth represents the methods of Twitter OAuth2 authorization.
type Auth interface {
	BuildAuthorizationURL() (string, *AuthorizationState)
	ExchangeCode(ctx context.Context, code, codeVerifier string) (*oauth2.Token, error)
}
