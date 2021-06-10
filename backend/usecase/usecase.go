package usecase

import "github.com/p1ass/midare/twitter"

type Usecase struct {
	twiCli twitter.Client
}

func NewUsecase(twiCli twitter.Client) *Usecase {
	return &Usecase{
		twiCli: twiCli,
	}
}
