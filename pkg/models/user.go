package models

import "gorm.io/gorm"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"size:100;not null"`
	LastName  string `gorm:"size:100;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
