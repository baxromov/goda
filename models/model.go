package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel defines common fields for all models.
type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"` // Primary key
	CreatedAt time.Time      `json:"created_at"`           // Timestamp for creation
	UpdatedAt time.Time      `json:"updated_at"`           // Timestamp for last update
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // Soft delete field
}
