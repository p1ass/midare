package datastore

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/mrjones/oauth"
)

func Test_client_AccessToken(t *testing.T) {
	fixed := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	now = func() time.Time {
		return fixed
	}

	type args struct {
		userID string
		aToken *oauth.AccessToken
	}
	tests := []struct {
		name           string
		args           args
		nowAfterStored time.Time
		want           *oauth.AccessToken
		wantStoreErr   bool
		wantFetchErr   bool
	}{
		{
			name: "保存したトークンを正しく取得できる",
			args: args{
				userID: uuid.NewString(),
				aToken: &oauth.AccessToken{
					Token:          "accessToken",
					Secret:         "requestSecret",
					AdditionalData: map[string]string{"screen_name": "p1ass"},
				},
			},
			nowAfterStored: fixed,
			want: &oauth.AccessToken{
				Token:          "accessToken",
				Secret:         "requestSecret",
				AdditionalData: map[string]string{"screen_name": "p1ass"},
			},
			wantStoreErr: false,
			wantFetchErr: false,
		},
		{
			name: "30分経過すると保存したトークンを取得できなくなる",
			args: args{
				userID: uuid.NewString(),
				aToken: &oauth.AccessToken{
					Token:  "accessToken",
					Secret: "requestSecret",
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

			if err := c.StoreAccessToken(context.Background(), tt.args.userID, tt.args.aToken); (err != nil) != tt.wantStoreErr {
				t.Errorf("StoreAccessToken() error = %v, wantErr %v", err, tt.wantStoreErr)
			}

			tmpNow := now
			now = func() time.Time {
				return tt.nowAfterStored
			}
			defer func() {
				now = tmpNow
			}()

			got, err := c.FetchAccessToken(context.Background(), tt.args.userID)
			if (err != nil) != tt.wantFetchErr {
				t.Errorf("FetchAccessToken() error = %v, wantErr %v", err, tt.wantFetchErr)
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("FetchAccessToken() got = %v, want %v diff= %v", got, tt.args.aToken, cmp.Diff(got, tt.want))
			}
		})
	}
}
