package handlers

import (
	"mfo-service/internal/logger"
	"mfo-service/internal/services"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "mfo-service/docs" // swagger docs
)

type (
	handlers struct {
		coldUsers *services.ColdUsersService
		l         logger.Logger
	}
)

func NewHandlers(
	coldUsers *services.ColdUsersService,
	log logger.Logger,
) *chi.Mux {
	r := chi.NewMux()

	h := handlers{
		coldUsers: coldUsers,
		l:         log,
	}

	h.build(r)

	return r
}

func (h *handlers) build(r chi.Router) {
	r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300,
    }))
	r.Use(middleware.Recoverer)
	r.Get("/docs/*", httpSwagger.WrapHandler)
	r.Get("/cold-users/{phone}", h.GetUserByPhone)
}
