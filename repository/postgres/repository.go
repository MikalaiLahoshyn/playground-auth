package postgres

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(options ...func(*Repository)) *Repository {
	repo := &Repository{}
	for _, option := range options {
		option(repo)
	}

	return repo
}

func WithDB(db *sqlx.DB) func(*Repository) {
	return func(r *Repository) {
		r.db = db
	}
}
