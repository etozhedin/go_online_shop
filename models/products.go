package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
	ImageURL    string `gorm:"type:text"` // URL or path to the product image
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
