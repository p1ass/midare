//go:build integration
// +build integration

package twitter

import (
	"os"
	"testing"

	"github.com/mrjones/oauth"
)

func Test_client_getUserTweets(t *testing.T) {

	token := &oauth.AccessToken{
		Token:          os.Getenv("TWITTER_ACCESS_TOKEN_FOR_TEST"),
		Secret:         os.Getenv("TWITTER_TOKEN_SECRET_FOR_TEST"),
		AdditionalData: nil,
	}

	tests := []struct {
		name       string
		token      *oauth.AccessToken
		screenName string
		want       []*Tweet
		wantErr    bool
	}{
		{
			name:       "p1ass",
			screenName: "p1ass",
			want:       nil,
			wantErr:    false,
		},
		{
			name:       "うじまる",
			screenName: "uzimaru0000",
			want:       nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := newClient(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"), "")
			_, err := cli.getUserTweets(token, tt.screenName, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("getUserTweets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
