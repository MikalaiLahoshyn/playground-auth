package handler

import (
	"auth/logging"
	"auth/models"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) Login(c echo.Context) error {
	var req models.LoginRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to decode request payload", logging.Any("error", err))
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Invalid input data", "bad payload": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		h.logger.Error("Failed to validate request payload", logging.Any("error", err))
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Validation failed", "validation error": err.Error()})
	}

	user, err := h.userService.GetUser(c.Request().Context(), req.Login)
	if err != nil {
		h.logger.Error("Failed to get user", logging.Any("error", err))
		return c.JSON(http.StatusNotFound, map[string]any{"message": "Failed to get user", "error": err.Error()})
	}

	err = h.userService.CheckCredentials(c.Request().Context(), *user, req.Password)

	if err != nil {
		h.logger.Error("Failed to check credentials", logging.Any("error", err))
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Failed to check creds", "error": err.Error()})
	}

	accessToken, refreshToken, err := h.tokenService.GenerateJWTTokenPair(c.Request().Context(), *user)
	if err != nil {
		h.logger.Error("Failed to generate token", logging.Any("error", err))
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Failed to generate token", "error": err.Error()})
	}

	// Set the refresh token in HTTP-only cookie
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   7 * 24 * 60 * 60, // 7 days
	})

	resp := models.LoginResponse{
		Token: accessToken,
	}

	return c.JSON(http.StatusOK, map[string]any{"message": "Login completed", "token": resp})
}
