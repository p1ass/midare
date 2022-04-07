package usecase

import (
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/p1ass/midare/datastore"
	"github.com/p1ass/midare/logging"
	"github.com/p1ass/midare/twitter"
	"golang.org/x/oauth2"
)

func TestUsecase_GetAwakePeriods_WhenUserFoundShouldReturnPeriodsWhichLengthIsOverZero(t *testing.T) {
	u := newUsecaseForTest(t)

	userID := "1032935958964973568"
	periods, _, err := u.GetAwakePeriods(newContextForTest(), userID, nil)

	if (err != nil) != false {
		t.Errorf("GetAwakePeriods() error = %v, wantErr %v", err, false)
		return
	}
	if len(periods) == 0 {
		t.Errorf("GetAwakePeriods() periods should have length over zero, but  %v", len(periods))
	}

}

func TestUsecase_GetAwakePeriods_WhenUserFoundShouldReturnShareURL(t *testing.T) {
	u := newUsecaseForTest(t)

	userID := "1032935958964973568"
	_, url, err := u.GetAwakePeriods(newContextForTest(), userID, nil)

	if (err != nil) != false {
		t.Errorf("GetAwakePeriods() error = %v, wantErr %v", err, false)
		return
	}

	if !strings.Contains(url, "http://localhost.local:3000/share/") {
		t.Errorf("GetAwakePeriods() url should contain share path, but not contain: got %v", url)
		return
	}

	shareID := strings.Replace(url, "http://localhost.local:3000/share/", "", 1)
	if shareID == "" {
		t.Errorf("GetAwakePeriods() url should contain shareID, but not contain: got %v", url)
		return
	}
}

func TestUsecase_AuthorizeToken(t *testing.T) {
	// TODO
}

func TestUsecase_GetLoginUrl(t *testing.T) {
	// TODO
}

func TestUsecase_GetUser_shouldReturnUser(t *testing.T) {
	u := newUsecaseForTest(t)

	user, err := u.GetUser(newContextForTest(), nil)

	if (err != nil) != false {
		t.Errorf("GetUser() error = %v, wantErr %v", err, false)
		return
	}

	if user == nil {
		t.Errorf("GetUser() user should not be nil but nil")
		return
	}
	if user.ID == "" {
		t.Errorf("GetUser() user.ID should not be empty but empty")
		return
	}
	if user.Name == "" {
		t.Errorf("GetUser() user.Name should not be empty but empty")
		return
	}
	if user.ScreenName == "" {
		t.Errorf("GetUser() user.ScreenName should not be empty but empty")
		return
	}
	if user.ImageURL == "" {
		t.Errorf("GetUser() user.ImageURL should not be empty but empty")
		return
	}
}

func TestUsecase_GetAccessToken_WhenAccessTokenFoundShouldReturnToken(t *testing.T) {
	u := newUsecaseForTest(t)

	ctx := newContextForTest()
	userID := "accessTokenFound"
	wantToken := &oauth2.Token{
		AccessToken: "dummyAccessToken",
	}

	err := u.dsCli.StoreAccessToken(ctx, userID, wantToken)
	if err != nil {
		t.Fatalf("StoreAccessToken() should return no error: but error = %v", err)
	}

	token, err := u.GetAccessToken(ctx, userID)

	if err != nil {
		t.Errorf("GetAccessToken() error = %v, wantErr %v", err, false)
		return
	}

	if !cmp.Equal(token, wantToken, cmpopts.IgnoreUnexported(oauth2.Token{})) {
		t.Errorf("GetAccessToken() token diff = %v", cmp.Diff(token, wantToken, cmpopts.IgnoreUnexported(oauth2.Token{})))
	}
}

func TestUsecase_GetAccessToken_WhenNoAccessTokenShouldReturnError(t *testing.T) {
	u := newUsecaseForTest(t)

	userID := "notFoundUserID"
	_, err := u.GetAccessToken(newContextForTest(), userID)

	if err == nil {
		t.Errorf("GetAccessToken() error should not be nil, but nil")
		return
	}
}

func newContextForTest() context.Context {
	return logging.Inject(context.Background(), logging.New())
}

func newUsecaseForTest(t *testing.T) *Usecase {
	t.Helper()

	// Inject fake client
	newTwitterClient = func(token *oauth2.Token) twitter.Client {
		return &twitter.FakeTwitterClient{}
	}

	dsCli, err := datastore.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	u := NewUsecase(nil, dsCli)
	return u
}
