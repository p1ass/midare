package main

import (
	"fmt"
	"log"
	"os"

	"github.com/p1ass/seikatsu-syukan-midare/twitter"
)

func main() {
	cli, err := twitter.NewTwitterClient(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"), "http://localhost.local:8080/callback")
	if err != nil {
		log.Fatalln(err)
	}
	url, err := cli.GetRequestTokenAndURL()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(url)
}
