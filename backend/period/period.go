package period

import (
	"time"

	"github.com/p1ass/midare/twitter"
)

const (
	// この時間以内にツイートされていたらその時間は起きていることにする
	awakeThreshold = 3*time.Hour + 30*time.Minute
)

type Period struct {
	OkiTime *twitter.Tweet `json:"okiTime"`
	NeTime  *twitter.Tweet `json:"neTime"`
}

// CalcAwakePeriods FIX ME テストで挙動を担保してはいるが、ロジックがブラックボックスなのでうまく整理したい
func CalcAwakePeriods(tweets []*twitter.Tweet) []*Period {
	periods := []*Period{}
	var neTweet *twitter.Tweet
	var okiTweet *twitter.Tweet
	var lastTweet *twitter.Tweet
	startIdx := 1
	for i, t := range tweets {
		if !t.ContainExcludedWord() {
			neTweet = t
			okiTweet = t
			lastTweet = t
			startIdx = i + 1
			break
		}
	}
	if lastTweet == nil {
		return periods
	}

	for _, t := range tweets[startIdx:] {
		if t.ContainExcludedWord() {
			continue
		}

		durationBetweenTweets := lastTweet.Created.Sub(t.Created)
		if durationBetweenTweets <= awakeThreshold {
			okiTweet = t
			lastTweet = t
			continue
		}

		if okiTweet != neTweet {
			periods = append(periods, &Period{
				OkiTime: okiTweet,
				NeTime:  neTweet,
			})
		}

		okiTweet = t
		neTweet = t
		lastTweet = t
	}

	if okiTweet != neTweet {
		periods = append(periods, &Period{
			OkiTime: okiTweet,
			NeTime:  neTweet,
		})
	}

	return periods
}
