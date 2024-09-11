package database

import (
	"go-homework/internal/entity"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("golang-homework.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB.AutoMigrate(&entity.Course{}, &entity.Teacher{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate the database: %v", err)
	}

}

func GetDB() *gorm.DB {
	return DB
}
