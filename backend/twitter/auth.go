package twitter

import (
	"context"
	"crypto/sha256"
	"encoding/base64"

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
type auth struct {
	config oauth2.Config
}

// AuthorizationState represents the state of the OAuth2 authorization state and PKCE code verifier.
type AuthorizationState struct {
	State        string
	CodeVerifier string
}

func NewAuth() Auth {
	cfg := config.NewTwitter()
	return &auth{
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

func (a *auth) BuildAuthorizationURL() (string, *AuthorizationState) {
	// for CSRF attack
	// https://datatracker.ietf.org/doc/html/rfc6749#section-10.12
	// SHOULD be less than or equal to 2^(-160) means 160bit / 8 = 20 bytes
	state := crypto.SecureRandomBase64Encoded(20)

	// Proof Key for Code Exchange (RFC 7636)
	// https://datatracker.ietf.org/doc/html/rfc7636
	codeVerifier := crypto.SecureRandomBase64Encoded(32)
	h := sha256.New()
	h.Write([]byte(codeVerifier))
	codeChallenge := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	url := a.config.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"))

	return url, &AuthorizationState{
		State:        state,
		CodeVerifier: codeVerifier,
	}
}

func (a *auth) ExchangeCode(ctx context.Context, code, codeVerifier string) (*oauth2.Token, error) {
	token, err := a.config.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	if err != nil {
		return nil, errors.NewForbidden("failed to exchange code: %v", err)
	}
	return token, nil
}
