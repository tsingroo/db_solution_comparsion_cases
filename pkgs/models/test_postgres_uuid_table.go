package models

import "github.com/google/uuid"

// TestPostgresUuidTable 对应 test_postgres_uuid_table 表
// 主键为应用层生成的 UUID（PostgreSQL 内置 UUID 类型）
type TestPostgresUuidTable struct {
	ID       uuid.UUID `gorm:"column:id;type:uuid;primaryKey"` // UUID 主键
	Name     string    `gorm:"column:name;type:varchar(50)"`   // 姓名
	Email    string    `gorm:"column:email;type:varchar(50)"`  // 邮箱
	Nickname string    `gorm:"column:nickname;type:varchar(50)"` // 昵称
}

// TableName 指定表名
func (TestPostgresUuidTable) TableName() string {
	return "test_postgres_uuid_table"
}
