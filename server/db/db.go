package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Link struct {
	ID        uint   `gorm:"index"`
	ShortLink string `gorm:"index"`
	LongLink  string
}

func Connect() {
	DB, _ = gorm.Open(postgres.Open(os.Getenv("DB_URL")))

	DB.AutoMigrate(&Link{})
}
