package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/p1ass/midare/datastore"
	"github.com/p1ass/midare/errors"
	"github.com/p1ass/midare/logging"
	"github.com/p1ass/midare/period"
	"github.com/p1ass/midare/twitter"
	"github.com/p1ass/midare/uploader"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// newTwitterClient uses for injecting fake client for unit test.
var newTwitterClient = twitter.NewClient

type Usecase struct {
	twiAuth       twitter.Auth
	dsCli         datastore.Client
	responseCache *cache.Cache
	imageUploader *uploader.ImageUploader
}

func NewUsecase(twiAuth twitter.Auth, dsCli datastore.Client) *Usecase {
	return &Usecase{
		twiAuth:       twiAuth,
		dsCli:         dsCli,
		responseCache: cache.New(5*time.Minute, 5*time.Minute),
		imageUploader: uploader.NewImageUploader(),
	}
}

// GetAwakePeriods gets awake periods using Twitter API.
func (u *Usecase) GetAwakePeriods(ctx context.Context, userID string, token *oauth2.Token) ([]*period.Period, string, error) {
	type getAwakePeriodsCache struct {
		Periods  []*period.Period `json:"periods"`
		ShareURL string           `json:"shareUrl"`
	}

	twiCli := newTwitterClient(token)

	cached, ok := u.responseCache.Get(userID)
	if ok {
		c := cached.(*getAwakePeriodsCache)
		return c.Periods, c.ShareURL, nil
	}

	tweets, err := twiCli.GetTweets(ctx, userID)
	if err != nil {
		return nil, "", err
	}

	periods := period.CalcAwakePeriods(tweets)

	shareID := uuid.New().String()

	url := u.imageUploader.Upload(ctx, periods, shareID, twiCli)

	res := &getAwakePeriodsCache{Periods: periods, ShareURL: url.String()}

	u.responseCache.SetDefault(userID, res)

	return periods, url.String(), nil
}

// AuthorizeToken exchanges code with access token.
// It is defined by OAuth2.
func (u *Usecase) AuthorizeToken(ctx context.Context, stateID, code, state string) (*twitter.User, error) {

	authState, err := u.dsCli.FetchAuthorizationState(ctx, stateID)
	if err != nil {
		if se, ok := errors.Cause(err).(*errors.ServiceError); ok && se.Code == errors.NotFound {
			return nil, errors.NewBadRequest("state not found: %s", stateID)
		}
		return nil, err
	}

	if state != authState.State {
		logging.Extract(ctx).Info("state not matched", zap.String("state", state), zap.String("expected", authState.State))
		return nil, errors.NewForbidden("state not matched")
	}

	token, err := u.twiAuth.ExchangeCode(ctx, code, authState.CodeVerifier)
	if err != nil {
		return nil, errors.Wrap(err, "exchange code")
	}

	user, err := u.GetUser(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "get user")
	}

	if err := u.dsCli.StoreAccessToken(ctx, user.ID, token); err != nil {
		return nil, errors.Wrap(err, "store access token")
	}
	return user, nil
}

// GetLoginUrl gets login url which starts OAuth2 flow.
// It is defined by OAuth2.
func (u *Usecase) GetLoginUrl(ctx context.Context, stateID string) (string, error) {
	url, authState := u.twiAuth.BuildAuthorizationURL()
	logging.Extract(ctx).Info(fmt.Sprintf("state id: %s", stateID), zap.String("state", authState.State))

	err := u.dsCli.StoreAuthorizationState(context.Background(), stateID, authState)
	if err != nil {
		return "", err
	}
	return url, nil
}

// GetUser gets user information using Twitter API.
func (u *Usecase) GetUser(ctx context.Context, token *oauth2.Token) (*twitter.User, error) {
	twiCli := newTwitterClient(token)

	user, err := twiCli.GetMe(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "account verify credentials")
	}
	return user, nil
}

// GetAccessToken gets access token from datastore.
func (u *Usecase) GetAccessToken(ctx context.Context, userID string) (*oauth2.Token, error) {
	token, err := u.dsCli.FetchAccessToken(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "fetch access token")
	}
	return token, nil
}
