package service

import (
	"auth/models"
	"auth/repository/postgres"
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenService struct {
	postgresRepo *postgres.Repository
}

func NewTokenService(postgresRepo *postgres.Repository) TokenService {
	return &tokenService{
		postgresRepo: postgresRepo,
	}
}

func (s *tokenService) GenerateJWTToken(ctx context.Context, user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_name": user.Name,
		"login":     user.Login,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), //add this one to config
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("hardcoded secret")) //add secret to config as well
	if err != nil {
		return "", fmt.Errorf("SERVICE-ERROR[%w]: failed to sign token: %s", models.ErrInternal, err.Error())
	}

	return signedToken, nil
}
