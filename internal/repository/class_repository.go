package repository

import (
	"context"
	"github.com/dijiacoder/go-web-learn/internal/model/gen/entity"
	"github.com/dijiacoder/go-web-learn/internal/model/gen/query"
	"gorm.io/gorm"
)

type ClassRepository struct {
	q  *query.Query
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{
		q:  query.Use(db),
		db: db,
	}
}

func (r *ClassRepository) ListAll(ctx context.Context) ([]*entity.Class, error) {
	classes, err := r.q.Class.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return classes, nil
}
