package twitter

import (
	"context"

	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/crypto"
	"github.com/p1ass/midare/errors"
	"golang.org/x/oauth2"
)

const (
	authorizationURL = "https://twitter.com/i/oauth2/authorize"
	tokenURL         = "https://api.twitter.com/2/oauth2/token"
)

// Auth represents the twitter OAuth2 authorization.
type Auth struct {
	config oauth2.Config
}

// AuthorizationState represents the state of the OAuth2 authorization state and PKCE code verifier.
type AuthorizationState struct {
	State        string
	CodeVerifier string
}

func NewAuth() *Auth {
	cfg := config.NewTwitter()
	return &Auth{
		config: oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Endpoint: oauth2.Endpoint{
				AuthURL:   authorizationURL,
				TokenURL:  tokenURL,
				AuthStyle: oauth2.AuthStyleInHeader,
			},
			RedirectURL: cfg.OAuthCallBackURL,
			Scopes:      []string{"tweet.read", "users.read"},
		},
	}
}

func (a *Auth) BuildAuthorizationURL() (string, *AuthorizationState) {
	state := crypto.ShortSecureRandomBase64()
	codeVerifier := crypto.ShortSecureRandomBase64()

	url := a.config.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge", codeVerifier),
		oauth2.SetAuthURLParam("code_challenge_method", "plain"))

	return url, &AuthorizationState{
		State:        state,
		CodeVerifier: codeVerifier,
	}
}

func (a *Auth) ExchangeCode(ctx context.Context, code, codeVerifier string) (*oauth2.Token, error) {
	token, err := a.config.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	if err != nil {
		return nil, errors.NewForbidden("failed to exchange code: %v", err)
	}
	return token, nil
}
