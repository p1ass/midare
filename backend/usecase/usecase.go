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

type getAwakePeriodsCache struct {
	Periods  []*period.Period `json:"periods"`
	ShareURL string           `json:"shareUrl"`
}
