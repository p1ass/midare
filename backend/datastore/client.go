package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/errors"
)

type Client interface {
	StoreRequestToken(ctx context.Context, rToken *oauth.RequestToken) error
	FetchRequestToken(ctx context.Context, token string) (*oauth.RequestToken, error)
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
