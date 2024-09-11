package usecase

import (
	"errors"
	"go-homework/internal/entity"
	"go-homework/internal/repository"

	"gorm.io/gorm"
)

type CourseService struct {
	CourseRepo *repository.CourseRepo
}

func NewCourseService(CourseRepo *repository.CourseRepo) *CourseService {
	return &CourseService{CourseRepo: CourseRepo}
}

func (s *CourseService) CreateCourse(course *entity.Course) error {
    existingCourse, err := s.CourseRepo.GetCourseByName(course.Name)

    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return err
    }

    if existingCourse != nil {
        return errors.New("course already exists")
    }

    return s.CourseRepo.CreateCourse(course)
}

func (s *CourseService) GetAllCourses() ([]entity.Course, error) {
	return s.CourseRepo.GetAllCourses()
}

func (s *CourseService) GetCourseByID(id uint) (*entity.Course, error) {
	return s.CourseRepo.GetCourseByID(id)
}

func (s *CourseService) UpdateCourse(id uint, updateCourse *entity.Course) error {
	return s.CourseRepo.UpdateCourse(id, updateCourse)
}

func (s *CourseService) DeleteCourse(id uint) error {
	return s.CourseRepo.DeleteCourse(id)
}
