// models/base.go
package models

import (
	"errors"
	"reflect"
	"time"
)

// Model is the base interface that all models should implement
type Model interface {
	Validate() error
	GetID() interface{}
}

// BaseModel provides common fields for all models
type BaseModel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetID returns the ID of the model
func (b *BaseModel) GetID() interface{} {
	return b.ID
}

// Validate performs basic validation on the model
func (b *BaseModel) Validate() error {
	return nil
}

// ValidateField is a helper function to validate a specific field using reflection
func ValidateField(model interface{}, fieldName string, validatorFn func(value interface{}) error) error {
	value := reflect.ValueOf(model)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	field := value.FieldByName(fieldName)
	if !field.IsValid() {
		return errors.New("field not found: " + fieldName)
	}

	return validatorFn(field.Interface())
}
