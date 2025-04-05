package main

import (
	"github.com/dijiacoder/go-web-learn/internal/api"
	"github.com/dijiacoder/go-web-learn/internal/config"
	"github.com/dijiacoder/go-web-learn/internal/query"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 初始化数据库
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("数据库初始化失败:", err)
	}

	// 初始化Gin
	r := gin.Default()

	// 创建Query实例
	q := query.Use(db)

	// 创建ClassHandler
	classHandler := api.NewClassHandler(q)

	// 设置路由
	r.POST("/classes", classHandler.CreateClass)
	r.GET("/classes/:id", classHandler.GetClass)
	r.PUT("/classes/:id", classHandler.UpdateClass)
	r.DELETE("/classes/:id", classHandler.DeleteClass)
	r.GET("/classes", classHandler.ListClasses)

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务失败:", err)
	}
}
