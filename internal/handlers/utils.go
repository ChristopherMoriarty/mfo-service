package handlers

import (
	"context"
	"errors"
	"net/http"

	"mfo-service/internal/apperrs"

	"github.com/go-chi/chi/v5"
)

func (h *handlers) handleError(ctx context.Context, w http.ResponseWriter, err error) {
	h.l.With("operation", chi.RouteContext(ctx).RoutePattern()).Error(err.Error())

	switch {
	case errors.Is(err, apperrs.ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, apperrs.ErrConditionViolation):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, apperrs.ErrAlreadyExist):
		w.WriteHeader(http.StatusConflict)
	case errors.Is(err, apperrs.ErrUnauthorize):
		w.WriteHeader(http.StatusUnauthorized)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, err = w.Write([]byte(err.Error()))
	if err != nil {
		h.l.Error("write error", err.Error())
		return
	}
}
