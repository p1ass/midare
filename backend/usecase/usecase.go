package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/period"
	"github.com/p1ass/midare/twitter"
	"github.com/p1ass/midare/uploader"
	"github.com/patrickmn/go-cache"
)

type Usecase struct {
	twiCli        twitter.Client
	responseCache *cache.Cache
	imageUploader *uploader.ImageUploader
}

func NewUsecase(twiCli twitter.Client) *Usecase {
	return &Usecase{
		twiCli:        twiCli,
		responseCache: cache.New(5*time.Minute, 5*time.Minute),
		imageUploader: uploader.NewImageUploader(twiCli),
	}
}

func (u *Usecase) GetAwakePeriods(accessToken *oauth.AccessToken) ([]*period.Period, string, error) {
	type getAwakePeriodsCache struct {
		Periods  []*period.Period `json:"periods"`
		ShareURL string           `json:"shareUrl"`
	}

	screenName := accessToken.AdditionalData["screen_name"]

	cached, ok := u.responseCache.Get(screenName)
	if ok {
		c := cached.(*getAwakePeriodsCache)
		return c.Periods, c.ShareURL, nil
	}

	tweets, err := u.twiCli.GetTweets(accessToken)
	if err != nil {
		return nil, "", err
	}

	periods := period.CalcAwakePeriods(tweets)

	shareID := uuid.New().String()

	url := u.imageUploader.Upload(periods, shareID, accessToken)

	res := &getAwakePeriodsCache{Periods: periods, ShareURL: url}

	u.responseCache.SetDefault(screenName, res)

	return periods, url, nil
}

func (u *Usecase) AuthorizeToken(token, verificationCode string) (*oauth.AccessToken, error) {
	accessToken, err := u.twiCli.AuthorizeToken(token, verificationCode)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (u *Usecase) GetLoginUrl() (string, error) {
	url, err := u.twiCli.GetLoginURL()
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *Usecase) GetUser(accessToken *oauth.AccessToken) (*twitter.TwitterUser, error) {
	user, err := u.twiCli.AccountVerifyCredentials(accessToken)
	if err != nil {
		return nil, err
	}
	return user, nil
}
