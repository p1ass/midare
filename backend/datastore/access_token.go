package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/p1ass/midare/errors"
	"golang.org/x/oauth2"
)

func (c client) StoreAccessToken(ctx context.Context, userID string, token *oauth2.Token) error {
	dto := &accessToken{
		Token:   token.AccessToken,
		Created: now(),
	}
	key := datastore.NameKey("OAuth2AccessToken", userID, nil)
	_, err := c.cli.Put(ctx, key, dto)
	if err != nil {
		return errors.Wrap(err, "failed to store access token")
	}
	return nil
}

func (c client) FetchAccessToken(ctx context.Context, userID string) (*oauth2.Token, error) {
	key := datastore.NameKey("OAuth2AccessToken", userID, nil)
	dto := &accessToken{}
	err := c.cli.Get(ctx, key, dto)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch access token")
	}

	// Redis時代と同様にセキュリティ上の理由から30分でタイムアウトするようにする
	if now().Sub(dto.Created) >= 30*time.Minute {
		err := c.cli.Delete(ctx, key)
		if err != nil {
			return nil, errors.Wrap(err, "failed to delete access token")
		}
		return nil, errors.New(errors.Unauthorized, "request token expired")
	}

	return &oauth2.Token{
		AccessToken: dto.Token,
	}, nil
}

type accessToken struct {
	Token   string
	Created time.Time
}
