package service

import (
	"auth/models"
	"auth/repository/postgres"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

type userService struct {
	postgresRepo *postgres.Repository
}

func NewUserService(postgresRepo *postgres.Repository) UserService {
	return &userService{
		postgresRepo: postgresRepo,
	}
}

func (s *userService) RegisterUser(ctx context.Context, user models.User) (int, error) {
	salt, err := generateSalt()
	if err != nil {
		return 0, fmt.Errorf("SERVICE-ERROR[%w]: failed to generate salt: %s", models.ErrInternal, err.Error())
	}

	hashedPassword, err := generateHash(user.Password, salt)
	if err != nil {
		return 0, fmt.Errorf("SERVICE-ERROR[%w]: failed to generate hash from password: %s", models.ErrInternal, err.Error())
	}

	user.Password = string(hashedPassword)

	id, err := s.postgresRepo.InsertUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("SERVICE-ERROR[%w]: failed to insert user: %s", models.ErrBadRequest, err.Error())
	}

	return id, nil
}

func (s *userService) GetUser(ctx context.Context, login string) (*models.User, error) {
	user, err := s.postgresRepo.GetUser(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("SERVICE-ERROR[%w]: failed to get user: %s", models.ErrNotFound, err.Error())
	}

	return user, nil
}

func (s *userService) CheckCredentials(ctx context.Context, user models.User, password string) error {
	//separate salt from actual hash
	data, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		return fmt.Errorf("SERVICE-ERROR[%w]: failed to decode stored hash: %s", models.ErrInternal, err)
	}

	if len(data) < saltLength {
		return fmt.Errorf("SERVICE-ERROR[%w]: invalid stored hash format: %s", models.ErrInternal, err)
	}

	salt := data[:saltLength]

	//generate hash using the same salt
	computedHash, err := generateHash(password, salt)
	if err != nil {
		return fmt.Errorf("SERVICE-ERROR[%w]: failed to generate hash: %s", models.ErrInternal, err)
	}

	if !bytes.Equal([]byte(computedHash), []byte(user.Password)) {
		return fmt.Errorf("SERVICE-ERROR[%w]: wrong password", models.ErrBadRequest)
	}

	return nil
}

const (
	saltLength  = 16        // Length of the salt in bytes
	timeCost    = 1         // Number of iterations
	memoryCost  = 64 * 1024 // Memory usage in KiB (64MB)
	parallelism = 4         // Number of threads
	keyLength   = 32        // Length of the generated key
)

func generateHash(input string, salt []byte) (string, error) {
	hash := argon2.IDKey([]byte(input), salt, timeCost, memoryCost, parallelism, keyLength)

	combined := append(salt, hash...)
	return base64.StdEncoding.EncodeToString(combined), nil
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	return salt, nil
}
