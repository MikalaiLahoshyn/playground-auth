package handler

import (
	"auth/models"
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) RegisterUser(c echo.Context) error {
	var req models.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Invalid input data"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Validation failed", "error": err})
	}

	user := models.InsertUser{
		Name:     req.Name,
		Surname:  req.Surname,
		Login:    req.Login,
		Password: req.Password,
	}

	id, err := h.userService.RegisterUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"error": "Failed to register user"})
	}

	response := models.RegisterResponse{
		ID:       id,
		Name:     user.Name,
		Surname:  user.Surname,
		Username: user.Login,
	}

	return c.JSON(http.StatusOK, map[string]any{"message": "User registered successfully", "response": response})
}
