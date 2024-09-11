package repository

import (
	course "go-homework/internal/entity"

	"gorm.io/gorm"
)

type ICourseRepository interface {
	CreateCourse(course *course.Course) error
	GetAllCourses() ([]course.Course, error)
	GetCourseByID(id uint) (*course.Course, error)
	UpdateCourse(id uint, course *course.Course) error
	DeleteCourse(id uint) error
	GetCourseByName(name string) (*course.Course, error)
}

type CourseRepo struct {
	db *gorm.DB
}

func NewCourseRepo(db *gorm.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

func (repo *CourseRepo) CreateCourse(course *course.Course) error {
	return repo.db.Create(course).Error
}

func (repo *CourseRepo) GetAllCourses() ([]course.Course, error) {
	var courses []course.Course
	err := repo.db.Preload("Courses").Find(&courses).Error
	return courses, err
}

func (repo *CourseRepo) GetCourseByID(id uint) (*course.Course, error) {
	var course course.Course
	err := repo.db.First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (repo *CourseRepo) UpdateCourse(id uint, updateCourse *course.Course) error {

	var course course.Course

	err := repo.db.First(&course, id).Error
	if err != nil {
		return err
	}

	err = repo.db.Model(&course).Updates(updateCourse).Error
	if err != nil {
		return err
	}

	return nil

}

func (repo *CourseRepo) DeleteCourse(id uint) error {
	return repo.db.Delete(&course.Course{}, id).Error
}

func (repo *CourseRepo) GetCourseByName(name string) (*course.Course, error) {
	var course course.Course
	err := repo.db.Where("name = ?", name).First(&course, name).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}
