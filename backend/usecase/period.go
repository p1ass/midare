package usecase

import (
	"strings"
	"time"

	"github.com/p1ass/midare/entity"
	"github.com/p1ass/midare/twitter"
)

const (
	// この時間以内にツイートされていたらその時間は起きていることにする
	awakeThreshold = 3*time.Hour + 30*time.Minute
)

func (u *Usecase) CalcAwakePeriods(ts []*twitter.Tweet) []*entity.Period {
	periods := []*entity.Period{}
	var neTweet *twitter.Tweet
	var okiTweet *twitter.Tweet
	var lastTweet *twitter.Tweet
	startIdx := 1
	for i, t := range ts {
		if !u.containExcludeWord(t.Text) {
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

	for _, t := range ts[startIdx:] {
		if u.containExcludeWord(t.Text) {
			continue
		}

		durationBetweenTweets := lastTweet.Created.Sub(t.Created)
		if durationBetweenTweets <= awakeThreshold {
			okiTweet = t
			lastTweet = t
			continue
		}

		if okiTweet != neTweet {
			periods = append(periods, &entity.Period{
				OkiTime: okiTweet,
				NeTime:  neTweet,
			})
		}

		okiTweet = t
		neTweet = t
		lastTweet = t
	}

	if okiTweet != neTweet {
		periods = append(periods, &entity.Period{
			OkiTime: okiTweet,
			NeTime:  neTweet,
		})
	}

	return periods
}

func (u *Usecase) containExcludeWord(text string) bool {
	excludeWords := []string{"ぼくへ 生活習慣乱れてませんか？", "みんなへ 生活習慣乱れてませんか？", "#contributter_report", "のポスト数"}
	for _, word := range excludeWords {
		if strings.Contains(text, word) {
			return true
		}
	}
	return false
}
