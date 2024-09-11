package usecase

import (
	"errors"
	"go-homework/internal/entity"
	"go-homework/internal/repository"

	"gorm.io/gorm"
)

type CourseService struct {
	courseRepo *repository.CourseRepo
}

func NewCourseService(courseRepo *repository.CourseRepo) *CourseService {
	return &CourseService{courseRepo: courseRepo}
}

func (s *CourseService) CreateCourse(course *entity.Course) error {
	existingCourse, err := s.courseRepo.GetCourseByName(course.Name)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existingCourse != nil {
		return errors.New("course already exists")
	}

	return s.courseRepo.CreateCourse(course)
}

func (s *CourseService) GetAllCourses() ([]entity.Course, error) {
	return s.courseRepo.GetAllCourses()
}

func (s *CourseService) GetCourseByID(id uint) (*entity.Course, error) {
	return s.courseRepo.GetCourseByID(id)
}

func (s *CourseService) UpdateCourse(id uint, updateCourse *entity.Course) error {
	return s.courseRepo.UpdateCourse(id, updateCourse)
}

func (s *CourseService) DeleteCourse(id uint) error {
	return s.courseRepo.DeleteCourse(id)
}
