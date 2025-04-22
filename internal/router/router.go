package router

import (
	"github.com/dijiacoder/go-web-learn/internal/handler"
	"github.com/dijiacoder/go-web-learn/internal/service"
	"github.com/gin-gonic/gin"
)

func NewRouter(container *service.Container) *gin.Engine {
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	studentHandler := handler.NewStudentHandler(container)
	classHandler := handler.NewClassHandler(container)
	r.GET("/student/:id", studentHandler.GetStudentByID)
	r.GET("/classes", classHandler.ListAll)
	return r
}
