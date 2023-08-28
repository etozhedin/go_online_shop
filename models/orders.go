package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Total     float64   `gorm:"not null"`
	Status    string    `gorm:"type:varchar(15);not null;default:'pending'"` // e.g., "pending", "shipped", "delivered"
	Products  []Product `gorm:"many2many:order_products;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
