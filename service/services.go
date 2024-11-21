package service

import (
	"auth/models"
	"context"
)

type OAuthService interface{}

type TwoFAService interface{}

type TokenService interface {
	GenerateJWTTokenPair(ctx context.Context, user models.User) (string, string, error)
}

type UserService interface {
	RegisterUser(ctx context.Context, user models.User) (int, error)
	GetUser(ctx context.Context, login string) (*models.User, error)
	CheckCredentials(ctx context.Context, user models.User, password string) error
}
