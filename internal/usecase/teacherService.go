package usecase

import (
	"errors"
	"go-homework/internal/entity"
	"go-homework/internal/repository"

	"gorm.io/gorm"
)

type TeacherService struct{
	teacherRepo *repository.TeacherRepo
}

func NewTeacherService(teacherRepo *repository.TeacherRepo) *TeacherService {
	return &TeacherService{teacherRepo: teacherRepo}
}

func (s *TeacherService) CreateTeacher(teacher *entity.Teacher) error {
	existingTeacher, err := s.teacherRepo.GetTeacherByPhone(teacher.Phone)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existingTeacher != nil {
		return errors.New("teacher already exists")
	}

	return s.teacherRepo.CreateTeacher(teacher)

}

func (s *TeacherService) GetTeacherByID(id uint) (*entity.Teacher, error) {
	return s.teacherRepo.GetTeacherByID(id)
}

func (s *TeacherService) GetAllTeachers() ([]entity.Teacher, error) {
	return s.teacherRepo.GetAllTeachers()
}

func (s *TeacherService) UpdateTeacher(id uint, updateTeacher *entity.Teacher) error {
	return s.teacherRepo.UpdateTeacher(id, updateTeacher)
}

func (s *TeacherService) DeleteTeacher(id uint) error {
	return s.teacherRepo.DeleteTeacher(id)
}