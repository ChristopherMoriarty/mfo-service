package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"mfo-service/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
