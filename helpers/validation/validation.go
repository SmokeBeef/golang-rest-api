package validation

import (
	"github.com/go-playground/validator/v10"
)

func Validate(field interface{}) error {
	validate := validator.Validate{}
	if err := validate.Struct(field); err != nil {
		return err
	}

	return nil
}
