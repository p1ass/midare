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

// CalcAwakePeriods calculates awake periods.
// IMPORTANT: For Twitter API specification, tweets slice is sorted by created_at desc.
func CalcAwakePeriods(tweets []*twitter.Tweet) []*Period {
	var periods []*Period

	neTweet, lastTweet, startIdx := getMostRecentValidTweet(tweets)

	// ツイートが全く無かった場合は空のスライスを返す
	if startIdx == 0 {
		return []*Period{}
	}

	for _, t := range tweets[startIdx:] {
		if t.ContainExcludedWord() {
			continue
		}

		durationBetweenTweets := lastTweet.Created.Sub(t.Created)
		if durationBetweenTweets > awakeThreshold {
			// しきい値時間を超えていればその時点でPeriodが確定するので、
			// lastTweetとneTweetが同じ場合を除きPeriodに追加する
			//(Periodが切り替わった後のtがしきい値以上間隔が空いている場合に発生)
			if lastTweet != neTweet {
				periods = append(periods, &Period{
					OkiTime: lastTweet,
					NeTime:  neTweet,
				})
			}
			// Periodが切り替わるので、neTweetを更新する
			neTweet = t
		}

		lastTweet = t
	}

	// ずっとしきい値以内だった場合はPeriodに追加されることなくループを抜けてしまうので、
	// ここでPeriodを追加する
	if lastTweet != neTweet {
		periods = append(periods, &Period{
			OkiTime: lastTweet,
			NeTime:  neTweet,
		})
	}

	// 0件の場合は、nullではなく空のJSONを返したいので
	if len(periods) == 0 {
		return []*Period{}
	}

	return periods
}

// 変数初期化のために、最も最近の除外されないツイートを探している
func getMostRecentValidTweet(tweets []*twitter.Tweet) (*twitter.Tweet, *twitter.Tweet, int) {
	var neTweet *twitter.Tweet
	var lastTweet *twitter.Tweet
	var startIdx int
	for i, t := range tweets {
		if !t.ContainExcludedWord() {
			neTweet = t
			lastTweet = t
			startIdx = i + 1
			break
		}
	}
	return neTweet, lastTweet, startIdx
}
