package validators

import "github.com/go-playground/validator"

func required(fl validator.FieldLevel) bool {
	return fl.Field().String() != ""
}
