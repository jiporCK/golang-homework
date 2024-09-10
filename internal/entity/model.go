package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Phone   string   `json:"phone"`
	Courses []Course `gorm:"foreignKey:TeacherID"`
}

type Course struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TeacherID   uint   `json:"teacher_id"`
}
