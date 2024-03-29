package datastore

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/p1ass/midare/errors"
	"github.com/p1ass/midare/logging"
	"github.com/p1ass/midare/twitter"
)

func Test_client_AuthorizationState(t *testing.T) {
	fixed := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	now = func() time.Time {
		return fixed
	}

	type args struct {
		stateID string
		state   *twitter.AuthorizationState
	}
	tests := []struct {
		name           string
		args           args
		nowAfterStored time.Time
		want           *twitter.AuthorizationState
		wantStoreErr   bool
		wantFetchErr   bool
	}{
		{
			name: "保存したトークンを正しく取得できる",
			args: args{
				stateID: uuid.NewString(),
				state:   &twitter.AuthorizationState{State: "state", CodeVerifier: "codeVerifier"},
			},
			nowAfterStored: fixed,
			want:           &twitter.AuthorizationState{State: "state", CodeVerifier: "codeVerifier"},
			wantStoreErr:   false,
			wantFetchErr:   false,
		},
		{
			name: "30分経過すると保存したトークンを取得できなくなる",
			args: args{
				stateID: uuid.NewString(),
				state:   &twitter.AuthorizationState{State: "state", CodeVerifier: "codeVerifier"},
			},
			nowAfterStored: fixed.Add(30 * time.Minute),
			want:           nil,
			wantStoreErr:   false,
			wantFetchErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient()
			if err != nil {
				t.Fatal(err)
			}

			ctx := context.Background()
			ctx = logging.Inject(ctx, logging.New())

			if err := c.StoreAuthorizationState(ctx, tt.args.stateID, tt.args.state); (err != nil) != tt.wantStoreErr {
				t.Errorf("StoreAccessToken() error = %v, wantErr %v", err, tt.wantStoreErr)
			}

			tmpNow := now
			now = func() time.Time {
				return tt.nowAfterStored
			}
			defer func() {
				now = tmpNow
			}()

			got, err := c.FetchAuthorizationState(ctx, tt.args.stateID)
			if (err != nil) != tt.wantFetchErr {
				t.Errorf("FetchAccessToken() error = %v, wantErr %v", err, tt.wantFetchErr)
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("FetchAuthorizationState() got = %v, want %v diff= %v", got, tt.args.state, cmp.Diff(got, tt.want))
			}
		})
	}
}

func Test_client_FetchAuthorizationStateShouldNotFoundErrorWhenNotFoundId(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	var notFoundID = "notFoundID"

	_, err = c.FetchAuthorizationState(context.Background(), notFoundID)
	se, ok := errors.Cause(err).(*errors.ServiceError)
	if !ok {
		t.Errorf("FetchAuthorizationState() error should ServiceError, but got %v", err)
		return
	}

	wantCode := errors.NotFound
	if se.Code != wantCode {
		t.Errorf("FetchAuthorizationState() errorCode = %v, wantErr %v", se.Code, wantCode)
	}
}
