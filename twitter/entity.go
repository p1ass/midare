package twitter

import "time"

// User represents a user info about twitter
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screenName"`
	ImageURL   string `json:"imageUrl"`
}

// Tweet represents a tweet
type Tweet struct {
	ID      string    `json:"id"`
	Text    string    `json:"text"`
	Created time.Time `json:"createdAt"`
}
