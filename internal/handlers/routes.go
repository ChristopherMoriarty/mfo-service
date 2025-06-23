package handlers

import (
	"mfo-service/internal/logger"
	"mfo-service/internal/services"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	r.Use(middleware.Recoverer)
	r.Get("/cold-users/{phone}", h.getUserByPhone)
}
