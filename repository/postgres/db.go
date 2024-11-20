package postgres

import (
	configs "auth/config"
	"auth/models"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns    = 10
	maxIdleConns    = 10
	maxConnLifetime = time.Minute * 3
)

// handlePostgresError handles postgres error and wraps it into corresponding entity error.
func handlePostgresError(name string, err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("[%s]: %w: %v", name, models.ErrNotFound, err.Error())
	}

	return fmt.Errorf("[%s]: %w: %v", name, models.ErrInternal, err.Error())
}

func OpenDB(cfg configs.PostgresDatabase) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	db, err := sqlx.Connect(cfg.Driver, connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL database: %w", err)
	}

	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	return db, nil
}