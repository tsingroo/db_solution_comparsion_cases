package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"db_optimization_techs/pkgs/dals"
	"db_optimization_techs/pkgs/models"
	"db_optimization_techs/pkgs/utils"

	"github.com/spf13/viper"
)

func main() {
	// 获取当前目录
	confPath := "."

	// 检查配置文件是否存在，不存在则创建
	configFile := filepath.Join(confPath, "config.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// 创建默认配置
		defaultConfig := models.Config{
			Database: models.DatabaseConfig{
				Type:     "mysql",
				Host:     "localhost",
				Port:     3306,
				User:     "root",
				Password: "",
				Database: "test_100m_db",
			},
		}
		configData, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			log.Fatalf("创建默认配置失败: %v", err)
		}
		if err := os.WriteFile(configFile, configData, 0644); err != nil {
			log.Fatalf("写入配置文件失败: %v", err)
		}
		log.Println("已创建默认配置文件:", configFile)
	}

	// 使用 viper 读取配置
	if err := utils.InitViper(confPath); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 解析配置到结构体
	var config models.Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("解析配置失败: %v", err)
	}

	// 初始化数据库连接
	db, err := dals.InitDB(&config.Database)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	log.Println("数据库连接成功")

	// 这里可以继续使用 db 进行数据库操作
	_ = db
}
