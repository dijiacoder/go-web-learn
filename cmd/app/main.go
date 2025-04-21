package main

import (
	"github.com/dijiacoder/go-web-learn/internal/config"
	"github.com/dijiacoder/go-web-learn/internal/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 初始化数据库
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("数据库初始化失败:", err)
	}

	//StudentHandler
	studentHandler := handler.NewStudentHandler(db)

	// 初始化Gin
	r := gin.Default()
	// 设置路由
	r.GET("/student/:id", studentHandler.GetStudentByID)

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务失败:", err)
	}
}
