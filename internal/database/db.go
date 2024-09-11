package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"log"
	"go-homework/internal/entity"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("golang-homework.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	db.AutoMigrate(&entity.Course{}, &entity.Teacher{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate the database: %v", err)
	}

}

func GetDB() *gorm.DB {
	return db
}

