package repository

import models "go-homework/internal/entity"

type TeacherRepository interface {
	CreateTeacher(teacher *models.Teacher) error
	GetTeacherByID(id uint) (*models.Teacher, error)
	GetAllTeachers() ([]models.Teacher, error)
	UpdateTeacher(teacher *models.Teacher) error
	DeleteTeacher(id uint) error
}

type CourseRepository interface {
	CreateCourse(course *models.Course) error
	GetCourseByID(id uint) (*models.Course, error)
	GetAllCourses() ([]models.Course, error)
	UpdateCourse(course *models.Course) error
	DeleteCourse(id uint) error
}
