package errors

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type ValidationError map[string]string

func NewValidationError(err error) ValidationError {
	validationError := ValidationError{}
	errors := err.(validator.ValidationErrors)

	for _, item := range(errors) {
		validationError[strings.ToLower(item.Field())] = tagMapping(item.Tag())
	}
	return validationError
}

func tagMapping(tag string) string {
	if tag == "required_without" {
		return "is required"
	}
	return tag
}
