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

	// 创建 PostgreSQL UUID DAL 实例
	dal := dals.NewTestPostgresUuidDAL(db)
	// 创建 PostgreSQL UUID Service 实例
	service := services.NewTestPostgresUuidService(dal)

	log.Println("开始性能测试...")

	createElapsed, err := service.Create()
	if err != nil {
		log.Fatalf("Create 失败: %v", err)
	}
	log.Printf("Create 完成，耗时: %d ms", createElapsed)

	getElapsed, err := service.Get()
	if err != nil {
		log.Fatalf("Get 失败: %v", err)
	}
	log.Printf("Get 完成，耗时: %d ms", getElapsed)

	updateElapsed, err := service.Update()
	if err != nil {
		log.Fatalf("Update 失败: %v", err)
	}
	log.Printf("Update 完成，耗时: %d ms", updateElapsed)

	deleteElapsed, err := service.Delete()
	if err != nil {
		log.Fatalf("Delete 失败: %v", err)
	}
	log.Printf("Delete 完成，耗时: %d ms", deleteElapsed)

	log.Println("性能测试完成")
}
