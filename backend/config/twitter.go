package config

import "os"

type Twitter struct {
	ConsumerKey      string
	ConsumerSecret   string
	OAuthCallBackURL string
}

func NewTwitter() *Twitter {
	return &Twitter{
		ConsumerKey:      os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret:   os.Getenv("TWITTER_CONSUMER_SECRET"),
		OAuthCallBackURL: os.Getenv("TWITTER_OAUTH_CALLBACK_URL"),
	}
}
