package repository

import (
	"go-homework/internal/entity"

	"gorm.io/gorm"
	"log"
)

type ICourseRepository interface {
	CreateCourse(course *entity.Course) error
	GetAllCourses() ([]entity.Course, error)
	GetCourseByID(id uint) (*entity.Course, error)
	UpdateCourse(id uint, course *entity.Course) error
	DeleteCourse(id uint) error
	GetCourseByName(name string) (*entity.Course, error)
}

type CourseRepo struct {
	DB *gorm.DB
}

func NewCourseRepo(db *gorm.DB) *CourseRepo {
	return &CourseRepo{DB: db}
}

func (repo *CourseRepo) CreateCourse(course *entity.Course) error {
    log.Printf("Creating course: %+v", course)
    return repo.DB.Create(course).Error
}

func (repo *CourseRepo) GetAllCourses() ([]entity.Course, error) {
    var courses []entity.Course
    err := repo.DB.Find(&courses).Error
    if err != nil {
        return nil, err
    }
    return courses, nil
}


func (repo *CourseRepo) GetCourseByID(id uint) (*entity.Course, error) {
	var course entity.Course
	err := repo.DB.First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (repo *CourseRepo) UpdateCourse(id uint, updateCourse *entity.Course) error {

	var course entity.Course

	err := repo.DB.First(&course, id).Error
	if err != nil {
		return err
	}

	err = repo.DB.Model(&course).Updates(updateCourse).Error
	if err != nil {
		return err
	}

	return nil

}

func (repo *CourseRepo) DeleteCourse(id uint) error {
	return repo.DB.Delete(&entity.Course{}, id).Error
}

func (repo *CourseRepo) GetCourseByName(name string) (*entity.Course, error) {
    var course entity.Course
    err := repo.DB.Where("name = ?", name).First(&course).Error
    if err != nil {
        return nil, err
    }
    return &course, nil
}

