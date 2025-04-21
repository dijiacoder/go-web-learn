package handler

import (
	"github.com/dijiacoder/go-web-learn/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type StudentHandler struct {
	studentService *service.StudentService
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{
		studentService: service.NewStudentService(db),
	}
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 ID"})
		return
	}

	student, err := h.studentService.GetStudentByID2(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}
