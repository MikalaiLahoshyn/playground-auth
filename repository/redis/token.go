package redis

import (
	"fmt"
	"time"
)

func (r *Repository) StoreRefreshToken(token string, userLogin string, ttl time.Duration) error {
	fail := func(err error) error {
		return handleRedisError("store refresh token", err)
	}

	refreshTokenKey := fmt.Sprintf("refresh_token:%s", userLogin)

	err := r.db.Set(refreshTokenKey, token, ttl).Err()
	if err != nil {
		return fail(fmt.Errorf("failed to store token: %w", err))
	}

	return nil
}
