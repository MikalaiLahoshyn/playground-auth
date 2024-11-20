package validators

import (
	"strconv"

	"github.com/go-playground/validator"
)

func max(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	param := fl.Param()
	num, _ := strconv.Atoi(param)
	return len(value) <= num
}
