package datastore

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/mrjones/oauth"
)

func Test_client_RequestToken(t *testing.T) {
	fixed := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	now = func() time.Time {
		return fixed
	}

	type args struct {
		rToken *oauth.RequestToken
	}
	tests := []struct {
		name           string
		args           args
		nowAfterStored time.Time
		want           *oauth.RequestToken
		wantStoreErr   bool
		wantFetchErr   bool
	}{
		{
			name: "保存したトークンを正しく取得できる",
			args: args{
				rToken: &oauth.RequestToken{
					Token:  "requestToken",
					Secret: "requestSecret",
				},
			},
			nowAfterStored: fixed,
			want: &oauth.RequestToken{
				Token:  "requestToken",
				Secret: "requestSecret",
			},
			wantStoreErr: false,
			wantFetchErr: false,
		},
		{
			name: "10分経過すると保存したトークンを取得できなくなる",
			args: args{
				rToken: &oauth.RequestToken{
					Token:  "requestToken",
					Secret: "requestSecret",
				},
			},
			nowAfterStored: fixed.Add(10 * time.Minute),
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

			if err := c.StoreRequestToken(context.Background(), tt.args.rToken); (err != nil) != tt.wantStoreErr {
				t.Errorf("StoreRequestToken() error = %v, wantErr %v", err, tt.wantStoreErr)
			}

			tmpNow := now
			now = func() time.Time {
				return tt.nowAfterStored
			}
			defer func() {
				now = tmpNow
			}()

			got, err := c.FetchRequestToken(context.Background(), tt.args.rToken.Token)
			if (err != nil) != tt.wantFetchErr {
				t.Errorf("FetchRequestToken() error = %v, wantErr %v", err, tt.wantFetchErr)
			}

			if !cmp.Equal(got, tt.want) {
				t.Errorf("FetchRequestToken() got = %v, want %v diff= %v", got, tt.args.rToken, cmp.Diff(got, tt.want))
			}
		})
	}
}
