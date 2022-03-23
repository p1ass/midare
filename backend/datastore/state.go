package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/p1ass/midare/errors"
	"github.com/p1ass/midare/twitter"
)

func (c client) StoreAuthorizationState(ctx context.Context, stateID string, authState *twitter.AuthorizationState) error {
	dto := &authorizationState{
		State:        authState.State,
		CodeVerifier: authState.CodeVerifier,
		Created:      now(),
	}
	key := datastore.NameKey("AuthorizationState", stateID, nil)
	_, err := c.cli.Put(ctx, key, dto)
	if err != nil {
		return errors.Wrap(err, "failed to store access token")
	}
	return nil
}

func (c client) FetchAuthorizationState(ctx context.Context, stateID string) (*twitter.AuthorizationState, error) {
	key := datastore.NameKey("AuthorizationState", stateID, nil)
	dto := &authorizationState{}
	err := c.cli.Get(ctx, key, dto)
	if err != nil {
		if errors.Cause(err) == datastore.ErrNoSuchEntity {
			return nil, errors.NewNotFound("state not found")
		}
		return nil, errors.Wrap(err, "failed to fetch authorization state")
	}

	if now().Sub(dto.Created) >= 15*time.Minute {
		err := c.cli.Delete(ctx, key)
		if err != nil {
			return nil, errors.Wrap(err, "failed to delete authorization state")
		}
		return nil, errors.NewNotFound("state not found")
	}

	return &twitter.AuthorizationState{
		State:        dto.State,
		CodeVerifier: dto.CodeVerifier,
	}, nil
}

type authorizationState struct {
	State        string
	CodeVerifier string
	Created      time.Time
}
