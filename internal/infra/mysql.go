package infra

import (
	"fmt"
	"github.com/dijiacoder/go-web-learn/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.Config) (*gorm.DB, error) {
	// 从配置文件中读取MySQL配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Mysql.User,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("连接MySQL失败: %v", err)
	}
	// 自动迁移schema
	//db.AutoMigrate(&model.User{})
	return db, nil
}
