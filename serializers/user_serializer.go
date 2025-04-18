package serializers

import (
	"errors"
	"goda/models"
)

// UserSerializer handles serialization and validation for the User model.
type UserSerializer struct {
	ModelSerializer
	Username string `json:"username" binding:"required,min=4"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty" binding:"required,min=6"`
}

// FromModel populates the serializer from a User model instance.
func (s *UserSerializer) FromModel(user models.User) {
	s.Username = user.Username
	s.Email = user.Email
	// Password is omitted intentionally for output
}

// ToModel converts the serializer back into a User model instance.
func (s *UserSerializer) ToModel() models.User {
	return models.User{
		Username: s.Username,
		Email:    s.Email,
		Password: s.Password,
	}
}

// Validate checks input data before creating or updating a user.
func (s *UserSerializer) Validate() error {
	// Validate struct fields
	if err := s.ValidateStruct(s); err != nil {
		return err
	}

	// Check for email uniqueness
	var count int64
	err := s.DB.Model(&models.User{}).Where("email = ?", s.Email).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("email is already in use")
	}

	return nil
}
