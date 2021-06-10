package usecase

import (
	"time"

	"github.com/p1ass/midare/twitter"
	"github.com/patrickmn/go-cache"
)

type Usecase struct {
	twiCli        twitter.Client
	responseCache *cache.Cache
}

func NewUsecase(twiCli twitter.Client) *Usecase {
	return &Usecase{
		twiCli:        twiCli,
		responseCache: cache.New(5*time.Minute, 5*time.Minute),
	}
}
