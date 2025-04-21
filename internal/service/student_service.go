package service

import (
	"context"
	"github.com/dijiacoder/go-web-learn/internal/models/entity"
	"github.com/dijiacoder/go-web-learn/internal/models/query"
	"gorm.io/gorm"
)

type StudentService struct {
	q  *query.Query
	db *gorm.DB
}

func NewStudentService(db *gorm.DB) *StudentService {
	return &StudentService{
		q:  query.Use(db),
		db: db,
	}
}

func (r *StudentService) GetStudentByID(id uint) (*entity.Student, error) {
	student, err := r.q.Student.WithContext(context.Background()).Where(r.q.Student.ID.Eq(int32(id))).First()
	return student, err
}

func (r *StudentService) GetStudentByID2(id uint) (*entity.Student, error) {
	var student entity.Student
	sql := "SELECT * FROM students WHERE id = ?"
	err := r.db.WithContext(context.Background()).Raw(sql, id).Scan(&student).Error
	return &student, err
}
