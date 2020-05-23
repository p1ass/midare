package twitter

// User represents a user info about twitter
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screenName"`
	ImageURL   string `json:"imageUrl"`
	UserID     string `json:"userID"`
}

// Credential represents a twitter oauth credential
type Credential struct {
	ID     int64
	token  string
	secret string
}

// NewCredential return a Credential
func NewCredential(id int64, token string, secret string) *Credential {
	return &Credential{ID: id, token: token, secret: secret}
}

// Token gets a oauth token
func (t *Credential) Token() string {
	return t.token
}

// Secret gets a oauth secret
func (t *Credential) Secret() string {
	return t.secret
}
