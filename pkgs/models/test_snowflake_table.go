package models

// TestSnowflakeTable 对应 test_snowflake_table 表
// 主键为 Sonyflake 算法生成的 int64
type TestSnowflakeTable struct {
	ID       int64  `gorm:"column:id;type:bigint;primaryKey"` // Sonyflake 主键
	Name     string `gorm:"column:name;type:varchar(50)"`     // 姓名
	Email    string `gorm:"column:email;type:varchar(50)"`    // 邮箱
	Nickname string `gorm:"column:nickname;type:varchar(50)"` // 昵称
}

// TableName 指定表名
func (TestSnowflakeTable) TableName() string {
	return "test_snowflake_table"
}
