package postgres

import (
	"auth/models"
	"context"
	"fmt"
)

func (r *Repository) InsertUser(ctx context.Context, user models.User) (int, error) {
	fail := func(err error) error {
		return handlePostgresError("insert user", err)
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, fail(fmt.Errorf("failed to begin a transaction: %w", err))
	}

	defer func() { _ = tx.Rollback() }()

	query := "INSERT INTO users (login, password_hash, name, surname) VALUES ($1, $2, $3, $4) RETURNING id;"

	var id int
	err = tx.QueryRowContext(ctx, query, user.Login, user.Password, user.Name, user.Surname).Scan(&id)
	if err != nil {
		return 0, fail(fmt.Errorf("failed to insert with context: %w", err))
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return id, nil
}

func (r *Repository) GetUser(ctx context.Context, login string) (*models.User, error) {
	fail := func(err error) error {
		return handlePostgresError("get user", err)
	}

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fail(fmt.Errorf("failed to begin a transaction: %w", err))
	}

	defer func() { _ = tx.Rollback() }()

	query := "SELECT name, surname, login, password_hash FROM users WHERE login = $1;"

	var user models.User
	err = tx.QueryRowContext(ctx, query, login).Scan(&user.Name, &user.Surname, &user.Login, &user.Password)
	if err != nil {
		return nil, fail(fmt.Errorf("failed to select with context: %w", err))
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &user, nil
}
