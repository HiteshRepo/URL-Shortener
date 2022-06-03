package main

import (
	"context"
	"github.com/hiteshpattanayak-tw/url_shortner/internal/app/di"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	app, err := di.InitializeApp(ctx)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	app.Start()
	<-interrupt()
	app.Shutdown()
}

func interrupt() chan os.Signal {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	return interrupt
}
