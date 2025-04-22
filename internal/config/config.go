package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Mysql MysqlConf `yaml:"mysql"`
}

type MysqlConf struct {
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func UnmarshalConfig(configFilePath string) (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(configFilePath)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &cfg, nil
}

func InitConfig() error {
	viper.SetConfigName("config")   // 配置文件名称（无扩展名）
	viper.SetConfigType("yaml")     // 配置文件类型
	viper.AddConfigPath("./config") // 查找配置文件的路径
	return viper.ReadInConfig()     // 读取配置文件
}
