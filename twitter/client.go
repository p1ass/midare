package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/mrjones/oauth"
	"github.com/p1ass/seikatsu-syukan-midare/errors"
)

// client is a client for twitter api
type client struct {
	consumer      *oauth.Consumer
	callbackURL   string
	requestTokens map[string]*oauth.RequestToken
	mu            sync.Mutex
}

// GetRequestTokenAndURL gets a request token associated with the user and login URL
func (cli *client) GetRequestTokenAndURL() (loginURL string, err error) {
	rToken, loginURL, err := cli.consumer.GetRequestTokenAndUrl(cli.callbackURL)
	if err != nil {
		return "", errors.Wrap(err, "failed to get access token")
	}

	cli.mu.Lock()
	cli.mu.Unlock()
	cli.requestTokens[rToken.Token] = rToken

	return loginURL, nil
}

// AuthorizeToken gets oauth access token
func (cli *client) AuthorizeToken(token, verificationCode string) (*Credential, error) {
	cli.mu.Lock()
	defer cli.mu.Unlock()
	rToken, ok := cli.requestTokens[token]
	if !ok {
		return nil, errors.New(errors.Unauthorized, "token secret not found")
	}

	aToken, err := cli.consumer.AuthorizeToken(rToken, verificationCode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to authorize token")
	}
	return NewCredential(0, aToken.Token, aToken.Secret), nil
}

// AccountVerifyCredentials fetch twitter profile from twitter api
func (cli *client) AccountVerifyCredentials(token, secret string) (*User, error) {
	httpCli, err := cli.httpClient(&oauth.AccessToken{
		Token:  token,
		Secret: secret,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to make http client")
	}

	resp, err := httpCli.Get(twitterAPIEndpoint + "/account/verify_credentials.json?include_email=true")
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch verify credentials from twitter api")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errMsg := &twitterError{}
		err = json.NewDecoder(resp.Body).Decode(errMsg)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode twitter api response")
		}
		return nil, errors.New(errors.Unauthorized, "twitter api response status code=%d message=%v", resp.StatusCode, errMsg.Errors)
	}

	res := &UserObject{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode twitter api response")
	}

	twiUser := &User{
		Name:       res.Name,
		ScreenName: res.ScreenName,
		ImageURL:   res.ProfileImageURL,
		UpdatedAt:  time.Now(),
	}
	twiUser = twiUser.SetID(res.ID)

	return twiUser, nil
}

// FriendsList fetch following list from twitter api
func (cli *client) FriendsList(cred Credential) ([]*UserObject, error) {
	type getFollowingResponse struct {
		Users          []*UserObject `json:"users"`
		NextCursor     int64         `json:"next_cursor"`
		PreviousCursor int64         `json:"previous_cursor"`
	}

	httpCli, err := cli.httpClient(&oauth.AccessToken{
		Token:  cred.Token(),
		Secret: cred.Secret(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to make http client")
	}

	query := url.Values{}
	query.Add("count", "200")
	query.Add("cursor", "-1")
	query.Add("skip_status", "true")
	query.Add("user_id", fmt.Sprintf("%d", cred.ID))

	resp, err := httpCli.Get(twitterAPIEndpoint + "/friends/list.json?" + query.Encode())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get following")

	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errMsg := &twitterError{}
		err = json.NewDecoder(resp.Body).Decode(errMsg)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode twitter api error response")
		}
		return nil, errors.New(errors.Unauthorized, "twitter api response status code=%d message=%v", resp.StatusCode, errMsg.Errors)
	}

	res := &getFollowingResponse{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode twitter api response")
	}

	return res.Users, nil
}

// GetUserByScreenName gets user which has argument screen name from Twitter API
func (cli *client) GetUserByScreenName(screenName string, cred Credential) (*UserObject, error) {
	httpCli, err := cli.httpClient(&oauth.AccessToken{
		Token:  cred.Token(),
		Secret: cred.Secret(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to make http client")
	}

	query := url.Values{}
	query.Add("screen_name", screenName)

	resp, err := httpCli.Get(twitterAPIEndpoint + "/users/lookup.json?" + query.Encode())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get following")

	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, errors.NewNotFound("user not found")
		}

		errMsg := &twitterError{}
		err = json.NewDecoder(resp.Body).Decode(errMsg)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode twitter api response")
		}
		return nil, errors.New(errors.Unauthorized, "twitter api response status code=%d message=%v", resp.StatusCode, errMsg.Errors)
	}

	res := []*UserObject{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode twitter api response")
	}

	if len(res) > 0 {
		return res[0], nil
	}

	return nil, errors.NewNotFound("user not found")

}

// httpClient make *http.Client using access token
func (cli *client) httpClient(aToken *oauth.AccessToken) (*http.Client, error) {
	client, err := cli.consumer.MakeHttpClient(aToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make http client from access token")
	}
	return client, nil
}

type twitterError struct {
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}
