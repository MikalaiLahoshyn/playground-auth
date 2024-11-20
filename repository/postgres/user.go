package postgres

import (
	"auth/models"
	"context"
	"fmt"
)

func (r *Repository) InsertUser(ctx context.Context, user models.InsertUser) (int, error) {
	fail := func(err error) error {
		return handlePostgresError("insert user", err)
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fail(fmt.Errorf("failed to begin a transaction: %w", err))
	}

	defer func() { _ = tx.Rollback() }()

	query := "INSERT INTO users(login, password_hash, name, surname) VALUES(?,?,?,?);"

	result, err := tx.ExecContext(ctx, query, user.Login, user.Password, user.Name, user.Surname)
	if err != nil {
		return 0, fail(fmt.Errorf("failed to insert with context: %w", err))
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fail(fmt.Errorf("failed to fetch inserted user's id: %w", err))
	}

	return int(id), nil
}
