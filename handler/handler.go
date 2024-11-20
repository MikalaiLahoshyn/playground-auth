package handler

import (
	"auth/logging"
	"auth/service"
)

type Handler struct {
	jwtService   service.JWTService
	oAuthService service.OAuthService
	twoFAService service.TwoFAService
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

func WithJWTServiceService(service service.JWTService) func(*Handler) {
	return func(handler *Handler) {
		handler.jwtService = service
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
