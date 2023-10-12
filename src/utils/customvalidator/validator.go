package customvalidator

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(v interface{}) error {
	return cv.Validator.Struct(v)
}
