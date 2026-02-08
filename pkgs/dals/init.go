package dals

import (
	"fmt"
	"log"
	"os"
	"time"

	"db_optimization_techs/pkgs/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 初始化数据库连接
// 根据配置创建 GORM 数据库连接并配置连接池
// 支持 MySQL 和 PostgreSQL
func InitDB(cfg *models.DatabaseConfig) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 2 * time.Second,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)

	var db *gorm.DB
	var err error

	// 根据数据库类型选择不同的驱动
	switch cfg.Type {
	case "mysql":
		// MySQL DSN 格式
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Database,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	case "postgresql":
		// PostgreSQL DSN 格式
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Password,
			cfg.Database,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	default:
		return nil, fmt.Errorf("不支持的数据库类型: %s", cfg.Type)
	}

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
