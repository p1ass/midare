package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/errors"
	"github.com/p1ass/midare/twitter"
	"golang.org/x/oauth2"
)

// テストで時間を差し替えられるように、変数としてnowを定義しておく
var now = time.Now

type Client interface {
	StoreAccessToken(ctx context.Context, userID string, token *oauth2.Token) error
	FetchAccessToken(ctx context.Context, userID string) (*oauth2.Token, error)

	StoreAuthorizationState(ctx context.Context, stateID string, authState *twitter.AuthorizationState) error
	FetchAuthorizationState(ctx context.Context, stateID string) (*twitter.AuthorizationState, error)
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
