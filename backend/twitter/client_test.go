package twitter

import (
	"net/http"
	"testing"

	"golang.org/x/oauth2"
)

func Test_authorizer_Add(t *testing.T) {
	t.Parallel()

	token := &oauth2.Token{
		AccessToken: "dummy_access_token",
		TokenType:   "Bearer",
	}

	a := &authorizer{
		token: token,
	}

	t.Run("Addを呼び出すごとで、HTTPリクエストのヘッダーにBearer Tokenが付与される", func(t *testing.T) {
		t.Parallel()

		req, _ := http.NewRequest("GET", "https://example.com", nil)

		a.Add(req)
		got := req.Header.Get("Authorization")

		want := "Bearer dummy_access_token"
		if want != got {
			t.Errorf("authorizer Add should set AUthorization Header, want %v, got %v", want, got)
		}
	})

}
