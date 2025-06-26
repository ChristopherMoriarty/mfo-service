package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"mfo-service/internal/app"
)

// @title MFO Service API
// @version 1.0
// @description API for working with cold users
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
