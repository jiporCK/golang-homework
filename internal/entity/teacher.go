package entity

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Phone   string   `json:"phone"`
	Courses []Course `gorm:"foreignKey:TeacherID"`
}


