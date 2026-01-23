package models

// Test100mTable 对应 test_100m_table 表
// 该表用于存储基本的用户信息
type Test100mTable struct {
	Uuid     string `gorm:"column:uuid;type:varchar(36)"`     // UUID 字符串
	Name     string `gorm:"column:name;type:varchar(50)"`     // 姓名
	Email    string `gorm:"column:email;type:varchar(50)"`    // 邮箱
	Nickname string `gorm:"column:nickname;type:varchar(50)"` // 昵称
}

// TableName 指定表名
func (Test100mTable) TableName() string {
	return "test_100m_table"
}
