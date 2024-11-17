package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
    Field string `json:"field"`
    Error string `json:"error"`
}

func GetErrMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return fe.Field() + " must be a valid email address"
	case "min":
		return fe.Field() + " must be at least " + fe.Param()
	case "max":
		return fe.Field() + " must not exceed " + fe.Param()
	default:
		return ""
	}
	
}

func GetErrValidation(err error) []ErrorMsg {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{fe.Field(), GetErrMsg(fe)}
		}
		return out
	}
	return nil
}
