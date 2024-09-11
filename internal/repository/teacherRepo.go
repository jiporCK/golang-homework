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
	DB *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) *TeacherRepo {
    return &TeacherRepo{DB: db}
}

func (repo *TeacherRepo) CreateTeacher(teacher *teacher.Teacher) error {
	return repo.DB.Create(teacher).Error
}

func (repo *TeacherRepo) GetAllTeachers() ([]teacher.Teacher, error) {
    var teachers []teacher.Teacher
    err := repo.DB.Preload("Courses").Find(&teachers).Error
    if err != nil {
        return nil, err
    }
    return teachers, nil
}



func (repo *TeacherRepo) GetTeacherByID(id uint) (*teacher.Teacher, error) {
	var teacher teacher.Teacher

	err := repo.DB.First(&teacher, id).Error
	if err != nil {
		return nil, err
	}

	return &teacher, nil
}

func (repo *TeacherRepo) UpdateTeacher(id uint, updateTeacher *teacher.Teacher) error {
	
	var teacher teacher.Teacher

	err := repo.DB.First(&teacher, id).Error
	if err != nil {
		return err
	}

	err = repo.DB.Model(teacher).Updates(updateTeacher).Error
	if err != nil {
		return err
	}

	return nil

}

func (repo *TeacherRepo) DeleteTeacher(id uint) error {
	return repo.DB.Delete(&teacher.Teacher{}, id).Error
}

func (repo *TeacherRepo) GetTeacherByPhone(phone string) (*teacher.Teacher, error) {
    var teacher teacher.Teacher
    err := repo.DB.Where("phone = ?", phone).First(&teacher).Error
    if err != nil {
        return nil, err
    }
    return &teacher, nil
}

