package models

import (
	"gorm.io/gorm"
)

// User defines a sample user model with fields and validation logic.
type User struct {
	BaseModel        // Embed BaseModel for common fields
	Username  string `gorm:"uniqueIndex" json:"username" binding:"required"` // Username must be unique
	Email     string `gorm:"uniqueIndex" json:"email" binding:"required"`    // Email must be unique
	Password  string `json:"-" binding:"required"`                           // Password (hidden in JSON responses)
}

// Validate is an example of a custom validation method for a User model.
func (u *User) Validate(db *gorm.DB) error {
	if len(u.Username) < 4 {
		return db.AddError(gorm.ErrInvalidData)
	}
	if len(u.Password) < 6 {
		return db.AddError(gorm.ErrInvalidData)
	}
	return nil
}
