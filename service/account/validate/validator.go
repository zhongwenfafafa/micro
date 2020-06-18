package validator

import (
	validator2 "github.com/go-playground/validator/v10"
)

func ValidatorRegister() error {
	validator := validator2.New()
}
