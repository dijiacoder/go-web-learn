package service

import (
	"context"
	"github.com/dijiacoder/go-web-learn/internal/model/gen/entity"
	"github.com/dijiacoder/go-web-learn/internal/repository"
)

type ClassService struct {
	classRepo *repository.ClassRepository
}

func NewClassService(container *repository.Container) *ClassService {
	return &ClassService{
		classRepo: container.ClassRepo,
	}
}

func (s *ClassService) ListAll(ctx context.Context) ([]*entity.Class, error) {
	classes, err := s.classRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	return classes, nil
}
