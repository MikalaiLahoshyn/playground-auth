package service

import (
	"auth/models"
	"auth/repository/postgres"
	"auth/repository/redis"
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenService struct {
	postgresRepo *postgres.Repository
	redisRepo    *redis.Repository
}

func NewTokenService(postgresRepo *postgres.Repository, redisRepo *redis.Repository) TokenService {
	return &tokenService{
		postgresRepo: postgresRepo,
		redisRepo:    redisRepo,
	}
}

func (s *tokenService) GenerateJWTTokenPair(ctx context.Context, user models.User) (string, string, error) {
	// Access Token - short TTL (e.g., 15 minutes)
	accessClaims := jwt.MapClaims{
		"user_name": user.Name,
		"login":     user.Login,
		"exp":       time.Now().Add(time.Minute * 15).Unix(), // Add TTL to config
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte("hardcoded access key secret"))
	if err != nil {
		return "", "", fmt.Errorf("SERVICE-ERROR[%w]: failed to sign access token: %s", models.ErrInternal, err.Error())
	}

	// Refresh Token - long TTL (e.g., 7 days)
	refreshClaims := jwt.MapClaims{
		"user_name": user.Name,
		"login":     user.Login,
		"exp":       time.Now().Add(time.Hour * 24 * 7).Unix(), // Add TTL to config
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("hardcoded refresh key secret"))
	if err != nil {
		return "", "", fmt.Errorf("SERVICE-ERROR[%w]: failed to sign refresh token: %s", models.ErrInternal, err.Error())
	}

	// TODO: Store the refresh token in Redis
	err = s.redisRepo.StoreRefreshToken(refreshTokenString, user.Login, time.Hour*24*7)
	if err != nil {
		return "", "", fmt.Errorf("SERVICE-ERROR[%w]: failed to store refresh token: %s", models.ErrInternal, err.Error())
	}

	// Return both tokens
	return accessTokenString, refreshTokenString, nil
}
