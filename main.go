package main

import (
	"fmt"
	"log"
	"os"

	"github.com/p1ass/seikatsu-syukan-midare/handler"

	"github.com/p1ass/seikatsu-syukan-midare/lib/logging"

	"github.com/p1ass/seikatsu-syukan-midare/twitter"
)

func main() {
	logger := logging.New()

	cli := twitter.NewClient(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"), "http://localhost.local:8080/callback")
	url, err := cli.GetRequestTokenAndURL()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(url)

	router, err := NewRouter(handler.NewHandler(cli, "http://localhost.local:3000", "http://localhost.local"), "*")
	if err != nil {
		log.Fatalln(err)
	}

	err = router.Run()
	if err != nil {
		logger.Panic("failed to listen and serve", logging.Error(err))
	}
}
