package databse

import (
	"fmt"
	"log"
	"online_shop/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto-migrate models
	if err := DB.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}
}
func GetDB() *gorm.DB {
	return DB
}
func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to retrieve underlying SQL database:", err)
	}
	sqlDB.Close()
}
