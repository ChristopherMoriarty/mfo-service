package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"mfo-service/internal/apperrs"
)

// GetUserByPhone godoc
// @Summary Get user by phone number
// @Description Get information about a cold user by phone number
// @Tags cold-users
// @Accept json
// @Produce json
// @Param phone path int true "User phone number"
// @Success 200 {object} domain.ColdUsers "User information"
// @Failure 400 {string} string "Invalid phone number"
// @Failure 500 {string} string "Internal server error"
// @Router /cold-users/{phone} [get]
func (h *handlers) GetUserByPhone(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	phoneStr := chi.URLParam(r, "phone")

	phone, err := strconv.Atoi(phoneStr)
	if err != nil {
		h.handleError(ctx, w, apperrs.ErrConditionViolation)
		return
	}

	user, err := h.coldUsers.GetColdUsers(ctx, phone)
	if err != nil {
		h.handleError(ctx, w, err)
		return
	}

	h.l.With("phone", phone).Info("Successfully found user")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		h.handleError(ctx, w, err)
	}
}

