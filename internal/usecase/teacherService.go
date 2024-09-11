package usecase

import (
	"errors"
	"go-homework/internal/entity"
	"go-homework/internal/repository"

	"gorm.io/gorm"
)

type TeacherService struct{
	TeacherRepo *repository.TeacherRepo
}

func NewTeacherService(TeacherRepo *repository.TeacherRepo) *TeacherService {
	return &TeacherService{TeacherRepo: TeacherRepo}
}

func (s *TeacherService) CreateTeacher(teacher *entity.Teacher) error {
    existingTeacher, err := s.TeacherRepo.GetTeacherByPhone(teacher.Phone)
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return err
    }

    if existingTeacher != nil {
        return errors.New("teacher already exists")
    }

    return s.TeacherRepo.CreateTeacher(teacher)
}


func (s *TeacherService) GetTeacherByID(id uint) (*entity.Teacher, error) {
	return s.TeacherRepo.GetTeacherByID(id)
}

func (s *TeacherService) GetAllTeachers() ([]entity.Teacher, error) {
	return s.TeacherRepo.GetAllTeachers()
}

func (s *TeacherService) UpdateTeacher(id uint, updateTeacher *entity.Teacher) error {
	return s.TeacherRepo.UpdateTeacher(id, updateTeacher)
}

func (s *TeacherService) DeleteTeacher(id uint) error {
	return s.TeacherRepo.DeleteTeacher(id)
}