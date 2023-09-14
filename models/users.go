package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint   `gorm:"primaryKey"`
	Username        string `gorm:"uniqueIndex;not null"`
	Email           string `gorm:"uniqueIndex;not null"`
	Password        string `gorm:"not null"`
	FullName        string
	ProfilePicURL   *string `gorm:"type:text"`
	Role            string  `gorm:"type:varchar(20);not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	PasswordExpiry  time.Time      `gorm:"not null"`
	PasswordRetries int            `gorm:"default:0"`
	IsActive        bool           `gorm:"default:true"`
}
