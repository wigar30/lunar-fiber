package model

import "github.com/go-playground/validator/v10"

type Validator struct {
	*validator.Validate
}
