package database

import (
	"log"

	"github.com/Magdiel-GVdz/accounts-manager/internal/domain"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("accounts.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&domain.User{})
	DB = db
	log.Println("Database connected successfully")
}
