package twitter

import (
	"context"
	"fmt"
	"time"

	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/p1ass/midare/errors"
	"github.com/p1ass/midare/logging"
	"go.uber.org/zap"
)

const (
	maxElapsedDuration = 21 * 24 * time.Hour

	tweetsCountPerAPI = 100
)

var (
	jst = time.FixedZone("Asia/Tokyo", 9*60*60)
)

func (c client) GetTweets(ctx context.Context, userID string) ([]*Tweet, error) {

	var mergedTweets []*Tweet
	paginationToken := ""
	// 一度のAPIで100件取得するので最大2000件になる
	for i := 0; i < 2000/tweetsCountPerAPI; i++ {
		opts := twitter.UserTweetTimelineOpts{
			TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt},
			MaxResults:  tweetsCountPerAPI,
		}
		if paginationToken != "" {
			opts.PaginationToken = paginationToken
		}
		res, err := c.cli.UserTweetTimeline(ctx, userID, opts)
		if err != nil {
			return nil, errors.Wrap(err, "fetch user tweet timeline")
		}
		tweets := toTweets(res.Raw.Tweets)

		logging.New().Info(fmt.Sprintf("rate limit: %d", res.RateLimit.Remaining), zap.Any("remaining", res.RateLimit.Remaining))

		if len(res.Raw.Tweets) == 0 {
			return []*Tweet{}, nil
		}
		extracted := extractWithinMaxElapsedDuration(tweets)
		mergedTweets = append(mergedTweets, extracted...)
		if doesReachFirstTweet(tweets) || exceededMaxElapsed(extracted, tweets) {
			break
		}
		paginationToken = res.Meta.NextToken
	}

	return mergedTweets, nil
}

func toTweets(tweetObjects []*twitter.TweetObj) []*Tweet {
	var ts []*Tweet

	for _, t := range tweetObjects {
		created, _ := time.Parse(time.RFC3339, t.CreatedAt)
		ts = append(ts, &Tweet{
			ID:      t.ID,
			Text:    t.Text,
			Created: created.In(jst),
		})
	}
	return ts
}

func exceededMaxElapsed(extracted, tweets []*Tweet) bool {
	return len(extracted) < len(tweets)
}

func doesReachFirstTweet(tweets []*Tweet) bool {
	return len(tweets) <= 1
}

func extractWithinMaxElapsedDuration(tweets []*Tweet) []*Tweet {
	var filtered []*Tweet

	for _, t := range tweets {
		if time.Since(t.Created) <= maxElapsedDuration {
			filtered = append(filtered, t)
		}
	}
	return filtered
}
