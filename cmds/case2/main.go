package main

import (
	"log"
	"os"
	"path/filepath"

	"db_optimization_techs/pkgs/dals"
	"db_optimization_techs/pkgs/models"
	"db_optimization_techs/pkgs/services"
	"db_optimization_techs/pkgs/utils"

	"github.com/spf13/viper"
)

func main() {
	// 从当前目录读取 config.json
	confPath := "."

	configFile := filepath.Join(confPath, "config.json")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatalf("配置文件不存在: %s，请先创建配置文件", configFile)
	}

	if err := utils.InitViper(confPath); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	var config models.Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("解析配置失败: %v", err)
	}

	db, err := dals.InitDB(&config.Database)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	log.Println("数据库连接成功")

	dal := dals.NewTest100mDAL(db)
	service := services.NewTest100mService(dal)

	duration, err := service.InsertBatch10000()
	if err != nil {
		log.Fatalf("批量插入 10000 条失败: %v", err)
	}
	log.Println("批量插入 10000 条成功，耗时:", duration, "ms")
}
