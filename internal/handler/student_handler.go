package handler

import (
	"context"
	"github.com/dijiacoder/go-web-learn/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	studentService *service.StudentService
}

func NewStudentHandler(container *service.Container) *StudentHandler {
	return &StudentHandler{
		studentService: container.StudentService,
	}
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 ID"})
		return
	}

	student, err := h.studentService.FindByID(context.Background(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}
