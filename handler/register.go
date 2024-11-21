package handler

import (
	"auth/logging"
	"auth/models"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) RegisterUser(c echo.Context) error {
	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Error("Failed to decode request payload", logging.Any("error", err))
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Invalid input data", "bad payload": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		h.logger.Error("Failed to validate request payload", logging.Any("error", err))
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Validation failed", "validation error": err.Error()})
	}

	user := &models.InsertUser{
		Name:     req.Name,
		Surname:  req.Surname,
		Login:    req.Login,
		Password: req.Password,
	}

	id, err := h.userService.RegisterUser(c.Request().Context(), *user)
	if err != nil {
		h.logger.Error("Failed to register user", logging.Any("error", err))
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Failed to register user"})
	}

	response := models.RegisterResponse{
		ID:       id,
		Name:     user.Name,
		Surname:  user.Surname,
		Username: user.Login,
	}

	return c.JSON(http.StatusOK, map[string]any{"message": "User registered successfully", "response": response})
}
