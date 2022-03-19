package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/errors"
)

func (c client) StoreAccessToken(ctx context.Context, userID string, aToken *oauth.AccessToken) error {
	dto := &accessToken{
		Token:      aToken.Token,
		Secret:     aToken.Secret,
		ScreenName: aToken.AdditionalData["screen_name"],
		Created:    now(),
	}
	key := datastore.NameKey("AccessToken", userID, nil)
	_, err := c.cli.Put(ctx, key, dto)
	if err != nil {
		return errors.Wrap(err, "failed to store access token")
	}
	return nil
}

func (c client) FetchAccessToken(ctx context.Context, userID string) (*oauth.AccessToken, error) {
	key := datastore.NameKey("AccessToken", userID, nil)
	dto := &accessToken{}
	err := c.cli.Get(ctx, key, dto)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch access token")
	}

	// Redis時代と同様にセキュリティ上の理由から30分でタイムアウトするようにする
	if now().Sub(dto.Created) >= 30*time.Minute {
		// TODO: ここでdatastoreから削除したいかも
		return nil, errors.New(errors.Unauthorized, "request token expired")
	}

	return &oauth.AccessToken{
		Token:          dto.Token,
		Secret:         dto.Secret,
		AdditionalData: map[string]string{"screen_name": dto.ScreenName},
	}, nil
}

type accessToken struct {
	Token      string
	Secret     string
	ScreenName string
	Created    time.Time
}
