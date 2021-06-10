package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/entity"
)

const (
	// この時間以内にツイートされていたらその時間は起きていることにする
	awakeThreshold = 3*time.Hour + 30*time.Minute
)

func (u *Usecase) GetAwakePeriods(accessToken *oauth.AccessToken) ([]*entity.Period, string, error) {
	screenName := accessToken.AdditionalData["screen_name"]

	cached, ok := u.responseCache.Get(screenName)
	if ok {
		c := cached.(*getAwakePeriodsCache)
		return c.Periods, c.ShareURL, nil
	}

	tweets, err := u.GetTweets(accessToken)
	if err != nil {
		return nil, "", err
	}

	periods := u.calcAwakePeriods(tweets)

	shareID := uuid.New().String()

	url := u.uploadImage(periods, shareID, accessToken)

	res := &getAwakePeriodsCache{Periods: periods, ShareURL: url}

	u.responseCache.SetDefault(screenName, res)

	return periods, url, nil
}

type getAwakePeriodsCache struct {
	Periods  []*entity.Period `json:"periods"`
	ShareURL string           `json:"shareUrl"`
}

// FIX ME テストで挙動を担保してはいるが、ロジックがブラックボックスなのでうまく整理したい
func (u *Usecase) calcAwakePeriods(ts []*entity.Tweet) []*entity.Period {
	periods := []*entity.Period{}
	var neTweet *entity.Tweet
	var okiTweet *entity.Tweet
	var lastTweet *entity.Tweet
	startIdx := 1
	for i, t := range ts {
		if !t.ContainExcludedWord() {
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
		if t.ContainExcludedWord() {
			continue
		}

		durationBetweenTweets := lastTweet.Created.Sub(t.Created)
		if durationBetweenTweets <= awakeThreshold {
			okiTweet = t
			lastTweet = t
			continue
		}

		if okiTweet != neTweet {
			periods = append(periods, &entity.Period{
				OkiTime: okiTweet,
				NeTime:  neTweet,
			})
		}

		okiTweet = t
		neTweet = t
		lastTweet = t
	}

	if okiTweet != neTweet {
		periods = append(periods, &entity.Period{
			OkiTime: okiTweet,
			NeTime:  neTweet,
		})
	}

	return periods
}
