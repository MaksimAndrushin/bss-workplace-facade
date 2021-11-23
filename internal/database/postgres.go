package database

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// NewPostgres returns DB
func NewPostgres(ctx context.Context, dsn string, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		log.Error().Msgf("failed to create database connection - %v", err)
		return nil, err
	}

	if err = db.PingContext(ctx); err != nil {
		log.Error().Msgf("failed ping the database - %v", err)
		return nil, err
	}

	return db, nil
}
