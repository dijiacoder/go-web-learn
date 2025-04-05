package api

import (
	"github.com/dijiacoder/go-web-learn/internal/model"
	"github.com/dijiacoder/go-web-learn/internal/query"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ClassHandler struct {
	q *query.Query
}

func NewClassHandler(q *query.Query) *ClassHandler {
	return &ClassHandler{q: q}
}

// CreateClass 创建班级
func (h *ClassHandler) CreateClass(c *gin.Context) {
	var class model.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.q.Class.WithContext(c.Request.Context()).Create(&class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, class)
}

// GetClass 获取班级
func (h *ClassHandler) GetClass(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	class, err := h.q.Class.WithContext(c.Request.Context()).Where(h.q.Class.ID.Eq(int32(id))).First()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	c.JSON(http.StatusOK, class)
}

// UpdateClass 更新班级
func (h *ClassHandler) UpdateClass(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.q.Class.WithContext(c.Request.Context()).Where(h.q.Class.ID.Eq(int32(id))).Updates(updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteClass 删除班级
func (h *ClassHandler) DeleteClass(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	_, err = h.q.Class.WithContext(c.Request.Context()).Where(h.q.Class.ID.Eq(int32(id))).Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// ListClasses 获取班级列表
func (h *ClassHandler) ListClasses(c *gin.Context) {
	classes, err := h.q.Class.WithContext(c.Request.Context()).Find()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, classes)
}
