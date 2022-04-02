package datastore

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/p1ass/midare/errors"
	"github.com/p1ass/midare/logging"
	"golang.org/x/oauth2"
)

func Test_client_AccessToken(t *testing.T) {
	fixed := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	now = func() time.Time {
		return fixed
	}

	type args struct {
		userID string
		token  *oauth2.Token
	}
	tests := []struct {
		name           string
		args           args
		nowAfterStored time.Time
		want           *oauth2.Token
		wantStoreErr   bool
		wantFetchErr   bool
	}{
		{
			name: "保存したトークンを正しく取得できる",
			args: args{
				userID: uuid.NewString(),
				token: &oauth2.Token{
					AccessToken: "accessToken",
				},
			},
			nowAfterStored: fixed,
			want: &oauth2.Token{
				AccessToken: "accessToken",
			},
			wantStoreErr: false,
			wantFetchErr: false,
		},
		{
			name: "30分経過すると保存したトークンを取得できなくなる",
			args: args{
				userID: uuid.NewString(),
				token: &oauth2.Token{
					AccessToken: "accessToken",
				},
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

			if err := c.StoreAccessToken(ctx, tt.args.userID, tt.args.token); (err != nil) != tt.wantStoreErr {
				t.Errorf("StoreAccessToken() error = %v, wantErr %v", err, tt.wantStoreErr)
			}

			tmpNow := now
			now = func() time.Time {
				return tt.nowAfterStored
			}
			defer func() {
				now = tmpNow
			}()

			got, err := c.FetchAccessToken(ctx, tt.args.userID)
			if (err != nil) != tt.wantFetchErr {
				t.Errorf("FetchAccessToken() error = %v, wantErr %v", err, tt.wantFetchErr)
			}

			if !cmp.Equal(got, tt.want, cmpopts.IgnoreUnexported(oauth2.Token{})) {
				t.Errorf("FetchAccessToken() got = %v, want %v diff= %v", got, tt.args.token, cmp.Diff(got, tt.want))
			}
		})
	}
}

func Test_client_FetchAccessTokenShouldNotFoundErrorWhenNotFoundId(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	var notFoundID = "notFoundID"

	_, err = c.FetchAccessToken(context.Background(), notFoundID)
	se, ok := errors.Cause(err).(*errors.ServiceError)
	if !ok {
		t.Errorf("FetchAccessToken() error should ServiceError, but got %v", err)
		return
	}

	wantCode := errors.NotFound
	if se.Code != wantCode {
		t.Errorf("FetchAccessToken() errorCode = %v, wantErr %v", se.Code, wantCode)
	}
}
