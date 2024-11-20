package validators

import (
	"regexp"

	"github.com/go-playground/validator"
)

func strongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	isValidLength := len(password) >= 8

	return hasLower && hasUpper && hasDigit && isValidLength
}
