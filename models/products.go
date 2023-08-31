package models

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	Price       float64
	ImageURL    string `gorm:"type:text"`
}
