package repository

import (
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
