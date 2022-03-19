package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/errors"
)

// テストで時間を差し替えられるように、変数としてnowを定義しておく
var now = time.Now

type Client interface {
	StoreRequestToken(ctx context.Context, rToken *oauth.RequestToken) error
	FetchRequestToken(ctx context.Context, token string) (*oauth.RequestToken, error)

	StoreAccessToken(ctx context.Context, userID string, aToken *oauth.AccessToken) error
	FetchAccessToken(ctx context.Context, userID string) (*oauth.AccessToken, error)
}

func NewClient() (Client, error) {
	ctx := context.Background()

	dsClient, err := datastore.NewClient(ctx, config.ReadDatastoreProjectId())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create datastore client")
	}

	return &client{
		cli: dsClient,
	}, nil
}

type client struct {
	cli *datastore.Client
}

var _ Client = &client{}
