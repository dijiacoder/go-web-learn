package handler

import (
	"context"
	"github.com/dijiacoder/go-web-learn/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClassHandler struct {
	classService *service.ClassService
}

func NewClassHandler(container *service.Container) *ClassHandler {
	return &ClassHandler{
		classService: container.ClassService,
	}
}

func (h *ClassHandler) ListAll(c *gin.Context) {

	classes, err := h.classService.ListAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, classes)
}
