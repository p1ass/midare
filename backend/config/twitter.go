package config

import "os"

type Twitter struct {
	ClientID         string
	ClientSecret     string
	OAuthCallBackURL string
}

func NewTwitter() *Twitter {
	return &Twitter{
		ClientID:         os.Getenv("TWITTER_CLIENT_ID"),
		ClientSecret:     os.Getenv("TWITTER_CLIENT_SECRET"),
		OAuthCallBackURL: os.Getenv("TWITTER_OAUTH_CALLBACK_URL"),
	}
}
