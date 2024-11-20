package service

import "auth/repository/postgres"

type oAuthService struct {
	postgresRepo *postgres.Repository
}

func NewOAuthService(postgresRepo *postgres.Repository) OAuthService {
	return &oAuthService{
		postgresRepo: postgresRepo,
	}
}
