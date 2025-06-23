package http

import (
	"database/sql"
	"net/http"

	"go.uber.org/zap"
)

type Handler struct {
	db  *sql.DB
	log *zap.Logger
}

func NewHandler(db *sql.DB, log *zap.Logger) *Handler {
	return &Handler{
		db:  db,
		log: log,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: implement your routes here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
