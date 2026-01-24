package dals

import (
	"fmt"
	"time"

	"db_optimization_techs/pkgs/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB 初始化数据库连接
// 根据配置创建 GORM 数据库连接并配置连接池
func InitDB(cfg *models.DatabaseConfig) (*gorm.DB, error) {
	// 构建 DSN 连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	// 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %w", err)
	}

	// 获取底层 *sql.DB 以配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("获取数据库连接失败: %w", err)
	}

	// 配置连接池参数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生存时间

	return db, nil
}
