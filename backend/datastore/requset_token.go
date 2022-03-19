package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/errors"
)

func (c client) StoreRequestToken(ctx context.Context, rToken *oauth.RequestToken) error {
	dto := &requestToken{
		Token:   rToken.Token,
		Secret:  rToken.Secret,
		Created: now(),
	}
	key := datastore.NameKey("RequestToken", rToken.Token, nil)
	_, err := c.cli.Put(ctx, key, dto)
	if err != nil {
		return errors.Wrap(err, "failed to store request token")
	}
	return nil
}

func (c client) FetchRequestToken(ctx context.Context, token string) (*oauth.RequestToken, error) {
	key := datastore.NameKey("RequestToken", token, nil)
	dto := &requestToken{}
	err := c.cli.Get(ctx, key, dto)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch request token")
	}

	// Redis時代と同様にセキュリティ上の理由から10分でタイムアウトするようにする
	if now().Sub(dto.Created) >= 10*time.Minute {
		return nil, errors.New(errors.Unauthorized, "request token expired")
	}

	return &oauth.RequestToken{
		Token:  dto.Token,
		Secret: dto.Secret,
	}, nil
}

type requestToken struct {
	Token   string
	Secret  string
	Created time.Time
}
