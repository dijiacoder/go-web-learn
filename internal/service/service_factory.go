package service

import (
	"github.com/dijiacoder/go-web-learn/internal/repository"
)

type Container struct {
	StudentService *StudentService
	ClassService   *ClassService

	//TODO
}

func NewServiceContainer(container *repository.Container) *Container {
	return &Container{
		StudentService: NewStudentService(container),
		ClassService:   NewClassService(container),
		//TODO
	}
}
