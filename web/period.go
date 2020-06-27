package web

import (
	"strings"

	"github.com/p1ass/midare/twitter"
)

func (h *Handler) calcAwakePeriods(ts []*twitter.Tweet) []*period {
	periods := []*period{}
	var neTweet *twitter.Tweet
	var okiTweet *twitter.Tweet
	var lastTweet *twitter.Tweet
	startIdx := 1
	for i, t := range ts {
		if !h.containExcludeWord(t.Text) {
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
		if h.containExcludeWord(t.Text) {
			continue
		}

		durationBetweenTweets := lastTweet.Created.Sub(t.Created)
		if durationBetweenTweets <= awakeThreshold {
			okiTweet = t
			lastTweet = t
			continue
		}

		if okiTweet != neTweet {
			periods = append(periods, &period{
				OkiTime: okiTweet,
				NeTime:  neTweet,
			})
		}

		okiTweet = t
		neTweet = t
		lastTweet = t
	}

	if okiTweet != neTweet {
		periods = append(periods, &period{
			OkiTime: okiTweet,
			NeTime:  neTweet,
		})
	}

	return periods
}

func (h *Handler) containExcludeWord(text string) bool {
	excludeWords := []string{"ぼくへ 生活習慣乱れてませんか？", "みんなへ 生活習慣乱れてませんか？", "#contributter_report", "のポスト数"}
	for _, word := range excludeWords {
		if strings.Contains(text, word) {
			return true
		}
	}
	return false
}

type period struct {
	OkiTime *twitter.Tweet `json:"okiTime"`
	NeTime  *twitter.Tweet `json:"neTime"`
}
