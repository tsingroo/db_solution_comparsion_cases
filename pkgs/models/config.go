package models

// DatabaseConfig 数据库配置结构体
// 支持 MySQL 和 PostgreSQL
type DatabaseConfig struct {
	Type     string `json:"type" mapstructure:"type"`         // 数据库类型: "mysql" 或 "postgresql"
	Host     string `json:"host" mapstructure:"host"`         // 数据库主机地址
	Port     int    `json:"port" mapstructure:"port"`         // 数据库端口
	User     string `json:"user" mapstructure:"user"`         // 数据库用户名
	Password string `json:"password" mapstructure:"password"` // 数据库密码
	Database string `json:"database" mapstructure:"database"` // 数据库名称
}

// Config 应用配置结构体
type Config struct {
	Database DatabaseConfig `json:"database" mapstructure:"database"` // 数据库配置
}
