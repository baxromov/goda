package serializers

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

// SerializerError holds detailed validation errors.
type SerializerError struct {
	Errors map[string]string `json:"errors"` // Field-level errors
}

// NewSerializerError creates a SerializerError for validation failures.
func NewSerializerError(err error) SerializerError {
	errors := map[string]string{}

	// Capture field-level error messages
	for _, validationErr := range err.(validator.ValidationErrors) {
		field := strings.ToLower(validationErr.Field())
		errors[field] = validationErr.Tag()
	}

	return SerializerError{Errors: errors}
}

// BaseSerializer serves as a generic base for all serializers.
type BaseSerializer struct {
	Validator *validator.Validate // Embedded validator instance
}

// ValidateStruct checks the struct against the defined validation rules.
func (s *BaseSerializer) ValidateStruct(obj interface{}) error {
	v := s.Validator
	if v == nil {
		v = validator.New() // If no validator is provided, create one
	}
	return v.Struct(obj)
}
