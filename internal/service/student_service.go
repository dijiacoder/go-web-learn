package service

import (
	"context"
	"github.com/dijiacoder/go-web-learn/internal/model/gen/entity"
	"github.com/dijiacoder/go-web-learn/internal/repository"
)

type StudentService struct {
	studentRepo *repository.StudentRepository
}

func NewStudentService(container *repository.Container) *StudentService {
	return &StudentService{
		studentRepo: container.StudentRepo,
	}
}

func (s *StudentService) FindByID(ctx context.Context, id int64) (*entity.Student, error) {
	student, err := s.studentRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return student, nil
}
