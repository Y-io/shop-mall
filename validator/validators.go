package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(field validator.FieldLevel) bool {
	mobile := field.Field().String()

	if ok, _ := regexp.MatchString(`^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`, mobile); !ok {
		return false
	}
	return true
}
