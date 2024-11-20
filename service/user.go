package service

import (
	"auth/models"
	"auth/repository/postgres"
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	postgresRepo *postgres.Repository
}

func NewUserService(postgresRepo *postgres.Repository) UserService {
	return &userService{
		postgresRepo: postgresRepo,
	}
}

func (s *userService) RegisterUser(ctx context.Context, user models.InsertUser) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("[%w]: failed to generate hash from passwrod", models.ErrInternal)
	}

	user.Password = string(hashedPassword)

	id, err := s.postgresRepo.InsertUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("[%w]: failed to insert user", models.ErrInternal)
	}

	return id, nil
}
