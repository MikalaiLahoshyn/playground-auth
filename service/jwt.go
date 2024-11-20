package service

import "auth/repository/postgres"

type jwtService struct {
	postgresRepo *postgres.Repository
}

func NewJWTService(postgresRepo *postgres.Repository) JWTService {
	return &jwtService{
		postgresRepo: postgresRepo,
	}
}
