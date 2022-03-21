package web

import (
	"github.com/p1ass/midare/datastore"
	"github.com/p1ass/midare/twitter"
	"github.com/p1ass/midare/usecase"
)

const (
	sevenDays = 60 * 60 * 24 * 7
)

// Handler is HTTP handler.
type Handler struct {
	frontendCallbackURL string
	dsCli               datastore.Client
	usecase             *usecase.Usecase
}

// NewHandler returns a new struct of Handler.
func NewHandler(twiAuth *twitter.Auth, dsCli datastore.Client, frontendCallbackURL string) (*Handler, error) {
	return &Handler{
		frontendCallbackURL: frontendCallbackURL,
		dsCli:               dsCli,
		usecase:             usecase.NewUsecase(twiAuth, dsCli),
	}, nil
}
