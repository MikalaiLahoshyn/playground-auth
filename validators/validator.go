package validators

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Return a friendly error message
		validationErrors := err.(validator.ValidationErrors)
		return errors.New(validationErrors.Error()) // Simplify for now
	}
	return nil
}

func RegisterValidators(e *echo.Echo) {
	validate := validator.New()

	_ = validate.RegisterValidation("strong_password", strongPassword)

	e.Validator = &CustomValidator{validator: validate}
}
