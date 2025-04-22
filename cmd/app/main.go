package main

import (
	"github.com/dijiacoder/go-web-learn/internal/config"
	"github.com/dijiacoder/go-web-learn/internal/infra"
	"github.com/dijiacoder/go-web-learn/internal/repository"
	router "github.com/dijiacoder/go-web-learn/internal/router"
	"github.com/dijiacoder/go-web-learn/internal/service"
	"log"
)

const (
	defaultConfigPath = "./configs/config.yaml"
)

func main() {

	cfg, err := config.UnmarshalConfig(defaultConfigPath)
	if err != nil {
		log.Fatal("配置文件解析失败:", err)
	}

	db, err := infra.InitMysql(cfg)
	if err != nil {
		log.Fatal("init mysql error", err)
	}

	repositoryContainer := repository.NewRepositoryContainer(db)
	serviceContainer := service.NewServiceContainer(repositoryContainer)
	r := router.NewRouter(serviceContainer)

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务失败:", err)
	}
	log.Println("服务启动成功:8080")
}
