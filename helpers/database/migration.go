package database

import (
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"gorm.io/gorm"
)

func DoMigration(db *gorm.DB) {
	db.AutoMigrate(domain.Product{})
}
