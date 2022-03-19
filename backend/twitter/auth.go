package twitter

import (
	"context"
	"encoding/json"

	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/errors"
	"github.com/p1ass/midare/logging"
)

// GetLoginURL gets login URL
func (cli *client) GetLoginURL() (loginURL string, err error) {
	rToken, loginURL, err := cli.consumer.GetRequestTokenAndUrl(cli.callbackURL)
	if err != nil {
		return "", errors.Wrap(err, "failed to get access token")
	}

	err = cli.dsCli.StoreRequestToken(context.Background(), rToken)
	if err != nil {
		logging.New().Error("failed to set request token to datastore", logging.Error(err))
		return "", err
	}

	return loginURL, nil
}

// AuthorizeToken gets oauth access token
func (cli *client) AuthorizeToken(token, verificationCode string) (*oauth.AccessToken, error) {

	rToken, err := cli.dsCli.FetchRequestToken(context.Background(), token)
	if err != nil {
		logging.New().Error("failed to get request token from datastore", logging.Error(err))
		return nil, err
	}

	aToken, err := cli.consumer.AuthorizeToken(rToken, verificationCode)
	if err != nil {
		return nil, errors.Wrap(err, "failed to authorize token")
	}

	return aToken, nil
}

// AccountVerifyCredentials fetch Twitter profile from Twitter api
func (cli *client) AccountVerifyCredentials(token *oauth.AccessToken) (*TwitterUser, error) {
	httpCli, err := cli.httpClient(token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make http client")
	}

	resp, err := httpCli.Get(twitterAPIEndpoint + "/account/verify_credentials.json")
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

	res := &userObject{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode twitter api response")
	}

	twiUser := &TwitterUser{
		ID:         res.IDStr,
		Name:       res.Name,
		ScreenName: res.ScreenName,
		ImageURL:   res.ProfileImageURL,
	}

	return twiUser, nil
}

// userObject is a user object for Twitter api
type userObject struct {
	ID              int64  `json:"id"`
	IDStr           string `json:"id_str"`
	Name            string `json:"name"`
	ScreenName      string `json:"screen_name"`
	URL             string `json:"url"`
	ProfileImageURL string `json:"profile_image_url_https"`
}
