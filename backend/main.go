package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/logging"
	"go.uber.org/zap"

	"cloud.google.com/go/profiler"
	"github.com/p1ass/midare/web"

	"github.com/p1ass/midare/twitter"
)

func main() {
	revision := config.ReadCloudRunRevision()
	if revision != "" {
		cfg := profiler.Config{
			Service:        "midare",
			ServiceVersion: revision,
			MutexProfiling: true,
		}
		if err := profiler.Start(cfg); err != nil {
			logging.New().Fatal("Profiler failed to start", zap.Error(err))
			return
		}
	}

	cli := twitter.NewClient()

	handler, err := web.NewHandler(cli, config.ReadFrontEndCallbackURL())
	if err != nil {
		logging.New().Fatal("Failed to initialize web handler", zap.Error(err))
		return
	}
	router, err := web.NewRouter(handler, config.ReadAllowCORSOriginURL())
	if err != nil {
		logging.New().Fatal("Failed to initialize web router", zap.Error(err))
		return
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.New().Fatal("Failed to listen and serve", zap.Error(err))
			return
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	logging.New().Info("Graceful Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logging.New().Fatal("Failed to shutdown server", zap.Error(err))
	}
	logging.New().Info("Server finished")

}
