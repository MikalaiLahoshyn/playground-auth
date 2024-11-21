package handler

import (
	"auth/logging"
	"auth/service"
)

type Handler struct {
	userService  service.UserService
	oAuthService service.OAuthService
	twoFAService service.TwoFAService
	tokenService service.TokenService
	logger       logging.Logger
}

func NewHandler(logger logging.Logger, options ...func(*Handler)) *Handler {
	handler := &Handler{
		logger: logger,
	}

	for _, option := range options {
		option(handler)
	}

	return handler
}

func WithUserService(service service.UserService) func(*Handler) {
	return func(handler *Handler) {
		handler.userService = service
	}
}

func WithOAuthService(service service.OAuthService) func(*Handler) {
	return func(handler *Handler) {
		handler.oAuthService = service
	}
}

func WithTwoFAService(service service.TwoFAService) func(*Handler) {
	return func(handler *Handler) {
		handler.twoFAService = service
	}
}

func WithTokenService(service service.TokenService) func(*Handler) {
	return func(handler *Handler) {
		handler.tokenService = service
	}
}
