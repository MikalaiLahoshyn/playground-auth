package redis

import (
	"github.com/go-redis/redis"
)

type Repository struct {
	db *redis.Client
}

func NewRepository(options ...func(*Repository)) *Repository {
	repo := &Repository{}
	for _, option := range options {
		option(repo)
	}

	return repo
}

func WithDB(db *redis.Client) func(*Repository) {
	return func(r *Repository) {
		r.db = db
	}
}
