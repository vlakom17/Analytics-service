package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/vlakom17/analytics-service/internal/config"
)

func NewPostgresConnection(cfg *config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
