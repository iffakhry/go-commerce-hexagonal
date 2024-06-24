package migrations

import (
	"github.com/iffakhry/go-commerce-hexagonal/internal/adapter/storages/postgres/models"
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Transaction{})
}
