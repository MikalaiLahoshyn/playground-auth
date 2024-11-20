package validators

import (
	"regexp"

	"github.com/go-playground/validator"
)

func email(fl validator.FieldLevel) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	email := fl.Field().String()
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
