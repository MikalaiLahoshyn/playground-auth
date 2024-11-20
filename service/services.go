package service

import (
	"auth/models"
	"context"
)

type OAuthService interface{}

type TwoFAService interface{}

type UserService interface {
	RegisterUser(ctx context.Context, user models.InsertUser) (int, error)
}
