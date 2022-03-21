package twitterv1

import (
	"strings"
	"time"
)

// Tweet represents a tweet
type Tweet struct {
	ID      string    `json:"id"`
	Text    string    `json:"text"`
	Created time.Time `json:"createdAt"`
}

func (t *Tweet) ContainExcludedWord() bool {
	excludeWords := []string{
		"ぼくへ 生活習慣乱れてませんか？",
		"みんなへ 生活習慣乱れてませんか？",
		"#contributter_report",
		"のポスト数",
	}
	for _, word := range excludeWords {
		if strings.Contains(t.Text, word) {
			return true
		}
	}
	return false
}

// TwitterUser represents a user info about twitter
type TwitterUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screenName"`
	ImageURL   string `json:"imageUrl"`
}
