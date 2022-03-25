package twitter

import (
	"crypto/sha256"
	"encoding/base64"
	"net/url"
	"testing"
)

func Test_auth_BuildAuthorizationURL(t *testing.T) {
	t.Parallel()

	a := NewAuth()

	t.Run("URLにはstateパラメータが含まれているべきである", func(t *testing.T) {
		t.Parallel()

		authUrl, state := a.BuildAuthorizationURL()

		parsedUrl, err := url.Parse(authUrl)
		if err != nil {
			t.Fatalf("authorization url should be valid url, but got url is %s and error: %v", authUrl, err)
		}

		gotState := parsedUrl.Query().Get("state")
		if gotState != state.State {
			t.Errorf("state %s should be included in authorization url, but got state is %s", state.State, gotState)
		}
	})

	t.Run("URLにはPKCE用のパラメータが含まれているべきである", func(t *testing.T) {
		t.Parallel()

		authUrl, state := a.BuildAuthorizationURL()

		parsedUrl, err := url.Parse(authUrl)
		if err != nil {
			t.Fatalf("authorization url should be valid url, but got url is %s and error: %v", authUrl, err)
		}

		wantCodeChallengeMethod := "S256"
		gotCodeChallengeMethod := parsedUrl.Query().Get("code_challenge_method")
		if wantCodeChallengeMethod != gotCodeChallengeMethod {
			t.Errorf("code_challenge_method %s should be included in authorization url, but got code_challenge_method is %s", wantCodeChallengeMethod, gotCodeChallengeMethod)
		}

		h := sha256.New()
		h.Write([]byte(state.CodeVerifier))
		wantCodeChallenge := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

		gotCodeChallenge := parsedUrl.Query().Get("code_challenge")
		if gotCodeChallenge != wantCodeChallenge {
			t.Errorf("code_challenge %s included inauthorization url should be hashed by sha256 and be encoded by base64 url encoded,"+
				"but got is %s", wantCodeChallenge, gotCodeChallenge)
		}
	})

}
