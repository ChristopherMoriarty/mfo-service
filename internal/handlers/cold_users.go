package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handlers) getUserByPhone(w http.ResponseWriter, r *http.Request) {
	phoneStr := chi.URLParam(r, "phone")
	phone, err := strconv.Atoi(phoneStr)
	if err != nil {
		http.Error(w, "Invalid phone number", http.StatusBadRequest)
		return
	}

	user, err := h.coldUsers.GetColdUsers(r.Context(), phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
