package repository

import (
	"gorm.io/gorm"
)

type Container struct {
	StudentRepo *StudentRepository
	ClassRepo   *ClassRepository

	//TODO
}

func NewRepositoryContainer(db *gorm.DB) *Container {
	return &Container{
		StudentRepo: NewStudentRepository(db),
		ClassRepo:   NewClassRepository(db),

		//TODO
	}
}
