package repository

import (
	"context"
	"github.com/dijiacoder/go-web-learn/internal/model/gen/entity"
	"github.com/dijiacoder/go-web-learn/internal/model/gen/query"
	"gorm.io/gorm"
)

type StudentRepository struct {
	q  *query.Query
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{
		q:  query.Use(db),
		db: db,
	}
}

func (r *StudentRepository) FindByID(ctx context.Context, id int64) (*entity.Student, error) {
	student, err := r.q.Student.WithContext(ctx).Where(r.q.Student.ID.Eq(int32(id))).First()
	return student, err
}

func (r *StudentRepository) Count(ctx context.Context) (count int64, err error) {
	return r.q.Student.WithContext(ctx).Count()
}

func (r *StudentRepository) ListByClassID(ctx context.Context, classID int64) (students []*entity.Student, err error) {
	sql := "SELECT * FROM students WHERE class_id = ?"
	err = r.db.WithContext(ctx).Raw(sql, classID).Scan(&students).Error
	return students, err
}
