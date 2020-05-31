package main

import (
	"log"
	"os"

	"github.com/p1ass/midare/web"

	"github.com/p1ass/midare/lib/logging"

	"github.com/p1ass/midare/twitter"
)

func main() {
	cli := twitter.NewClient(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"), os.Getenv("TWITTER_OAUTH_CALLBACK_URL"))

	router, err := web.NewRouter(web.NewHandler(cli, os.Getenv("FRONTEND_CALLBACK_URL")), os.Getenv("CORS_ALLOW_ORIGIN"))
	if err != nil {
		log.Fatalln(err)
	}

	err = router.Run()
	if err != nil {
		logger := logging.New()
		logger.Panic("failed to listen and serve", logging.Error(err))
	}
}
