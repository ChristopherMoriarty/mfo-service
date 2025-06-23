package app

import (
	"context"
	"fmt"
	"time"

	"mfo-service/internal/config"
	"mfo-service/internal/databases"
	"mfo-service/internal/handlers"
	"mfo-service/internal/logger"
	"mfo-service/internal/repositories"
	"mfo-service/internal/services"
	"mfo-service/internal/transport/http"
)

const serverShutdownTimeout = 1 * time.Minute

func Run(ctx context.Context) error {
	cfg, err := config.Parse()
	if err != nil {
		return fmt.Errorf("build config: %w", err)
	}
	fmt.Println("DB URI:", cfg.DB.GetURI())
	fmt.Println("HTTP Port:", cfg.HTTP.Port)

	log, err := logger.New(&cfg.Logger)
	if err != nil {
		return fmt.Errorf("create logger: %w", err)
	}

	db, closeDB, err := databases.NewDB(&cfg.DB)
	if err != nil {
		return fmt.Errorf("create database: %w", err)
	}
	defer func() {
		if err := closeDB(); err != nil {
			log.Error("close db connection: ", err.Error())
		}
	}()

	coldUsersRepo := repositories.NewColdUsersRepository(db)
	coldUsersService := services.NewColdUsersService(coldUsersRepo)

	hs := handlers.NewHandlers(coldUsersService, log)
	stopHTTPServer, err := http.ServeHTTP(&cfg.HTTP, hs)
	if err != nil {
		return fmt.Errorf("start HTTP server: %w", err)
	}
	defer func() {
		log.Info("shutting down HTTP server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
		defer cancel()

		if err := stopHTTPServer(shutdownCtx); err != nil {
			log.Error("shutdown HTTP server: ", err.Error())
		} else {
			log.Info("HTTP server shutdown complete")
		}
	}()

	log.Infof("app started on port: %d", cfg.HTTP.Port)
	<-ctx.Done()

	return nil
}
