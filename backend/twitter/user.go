package twitter

import (
	"context"

	"github.com/g8rswimmer/go-twitter/v2"
	"github.com/p1ass/midare/errors"
)

func (c client) GetMe(ctx context.Context) (*User, error) {
	res, err := c.cli.AuthUserLookup(ctx, twitter.UserLookupOpts{
		UserFields: []twitter.UserField{
			twitter.UserFieldProfileImageURL,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "twitter api: auth user lookup")
	}
	return &User{
		ID:         res.Raw.Users[0].ID,
		Name:       res.Raw.Users[0].Name,
		ScreenName: res.Raw.Users[0].UserName,
		ImageURL:   res.Raw.Users[0].ProfileImageURL,
	}, nil
}
