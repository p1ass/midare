package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"cloud.google.com/go/profiler"
	"github.com/p1ass/midare/web"

	"github.com/p1ass/midare/twitter"
)

func main() {
	if os.Getenv("K_REVISION") != "" {
		cfg := profiler.Config{
			Service:        "midare",
			ServiceVersion: os.Getenv("K_REVISION"),
			MutexProfiling: true,
		}
		if err := profiler.Start(cfg); err != nil {
			log.Fatalf("Profiler failed to start: %v", err)
		}
	}

	cli := twitter.NewClient(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"), os.Getenv("TWITTER_OAUTH_CALLBACK_URL"))

	handler, err := web.NewHandler(cli, os.Getenv("FRONTEND_CALLBACK_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	router, err := web.NewRouter(handler, os.Getenv("CORS_ALLOW_ORIGIN"))
	if err != nil {
		log.Fatalln(err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
