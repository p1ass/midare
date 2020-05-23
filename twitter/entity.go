package twitter

import (
	"fmt"
	"time"
)

// User represents a user info about twitter
type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	ScreenName string    `json:"screenName"`
	ImageURL   string    `json:"imageUrl"`
	UserID     string    `json:"userID"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// SetID sets id into ID and IDStr
func (u *User) SetID(id int64) *User {
	u.ID = fmt.Sprintf("%d", id)
	return u
}

// SetUserID sets user id into UserID and UserIDStr
func (u *User) SetUserID(userID int64) *User {
	u.UserID = fmt.Sprintf("%d", userID)
	return u
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
