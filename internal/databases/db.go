package databases

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"

	"mfo-service/internal/config"
)

type (
	closeFn func() error
	DB      = sql.DB
)

func NewDB(cfg *config.DB) (*sql.DB, closeFn, error) {
	connCfg, err := pgx.ParseURI(cfg.GetURI())
	if err != nil {
		return nil, nil, fmt.Errorf("parse URI: %w", err)
	}
	db := stdlib.OpenDB(connCfg)
	err = db.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("check connection: %w", err)
	}

	return db, db.Close, nil
}
