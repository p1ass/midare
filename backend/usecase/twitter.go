package usecase

import (
	"time"

	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/entity"
)

const (
	oldestTweetTime = 21 * 24 * time.Hour
)

// getTweets gets more than 2000 tweets.
func (u *Usecase) GetTweets(accessToken *oauth.AccessToken) ([]*entity.Tweet, error) {
	screenName := accessToken.AdditionalData["screen_name"]

	var allTweets []*entity.Tweet
	maxID := ""
	// 一度のAPIで200件取得するので最大2000件になる
	for i := 0; i < 10; i++ {
		tweets, err := u.twiCli.GetUserTweets(accessToken, screenName, maxID)
		if err != nil {
			return nil, err
		}
		if len(tweets) == 0 {
			return []*entity.Tweet{}, nil
		}
		filtered := u.filterByCreated(tweets)
		allTweets = append(allTweets, filtered...)
		if u.doesReachFirstTweet(tweets) || u.overOldestTweetTime(filtered, tweets) {
			break
		}
		maxID = allTweets[len(allTweets)-1].ID
	}

	return allTweets, nil
}

func (u *Usecase) overOldestTweetTime(filtered, tweets []*entity.Tweet) bool {
	return len(filtered) < len(tweets)
}

func (u *Usecase) doesReachFirstTweet(tweets []*entity.Tweet) bool {
	return len(tweets) <= 1
}

func (u *Usecase) filterByCreated(tweets []*entity.Tweet) []*entity.Tweet {
	var filtered []*entity.Tweet

	for _, t := range tweets {
		if time.Since(t.Created) <= oldestTweetTime {
			filtered = append(filtered, t)
		}
	}
	return filtered
}
