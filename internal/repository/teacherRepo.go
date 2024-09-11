package repository

import (
	teacher "go-homework/internal/entity"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	CreateTeacher(teacher *teacher.Teacher) error
	GetAllTeachers() ([]teacher.Teacher, error)
	GetTeacherByID(id uint) (*teacher.Teacher, error)
	UpdateTeacher(id uint, teacher *teacher.Teacher) error
	DeleteTeacher(id uint) error
}

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) *TeacherRepo {
	return &TeacherRepo{db: db}
}

func (repo *TeacherRepo) CreateTeacher(teacher *teacher.Teacher) error {
	return repo.db.Create(teacher).Error
}

func (repo *TeacherRepo) GetAllTeachers() ([]teacher.Teacher, error) {
	var teachers []teacher.Teacher

	err := repo.db.Preload("Teachers").Find(&teachers).Error
	return teachers, err
}

func (repo *TeacherRepo) GetTeacherByID(id uint) (*teacher.Teacher, error) {
	var teacher teacher.Teacher

	err := repo.db.First(&teacher, id).Error
	if err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (repo *TeacherRepo) UpdateTeacher(id uint, updateTeacher *teacher.Teacher) error {
	
	var teacher teacher.Teacher

	err := repo.db.First(&teacher, id).Error
	if err != nil {
		return err
	}

	err = repo.db.Model(teacher).Updates(updateTeacher).Error
	if err != nil {
		return err
	}

	return nil

}

func (repo *TeacherRepo) DeleteTeacher(id uint) error {
	return repo.db.Delete(&teacher.Teacher{}, id).Error
}

func (repo *TeacherRepo) GetTeacherByPhone(phone string) (*teacher.Teacher, error) {
	var teacher teacher.Teacher
	err := repo.db.Where("phone = ?", phone).First(&teacher, phone).Error
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}
