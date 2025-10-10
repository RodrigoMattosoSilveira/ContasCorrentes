package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
