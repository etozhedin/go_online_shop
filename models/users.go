package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint   `gorm:"primaryKey"`
	Username      string `gorm:"uniqueIndex;not null"`
	Email         string `gorm:"uniqueIndex;not null"`
	Password      string `gorm:"not null"` // Ensure this is hashed before storing!
	FullName      string
	ProfilePicURL *string `gorm:"type:text"` // Pointer to string to make it nullable
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
