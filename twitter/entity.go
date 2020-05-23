package twitter

// User represents a user info about twitter
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screenName"`
	ImageURL   string `json:"imageUrl"`
}
