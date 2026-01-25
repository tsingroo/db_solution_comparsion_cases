package models

// Test100mCrc32Table 对应 test_100m_crc32_table 表
// 该表包含 UUID 的 CRC32 值用于优化查询性能
type Test100mCrc32Table struct {
	UuidCrc32 uint32 `gorm:"column:uuid_crc32;type:int unsigned;primaryKey"` // UUID 的 CRC32 值，联合主键
	Uuid      string `gorm:"column:uuid;type:varchar(36);primaryKey"`        // UUID 字符串，联合主键
	Name      string `gorm:"column:name;type:varchar(50)"`                    // 姓名
	Email     string `gorm:"column:email;type:varchar(50)"`                    // 邮箱
	Nickname  string `gorm:"column:nickname;type:varchar(50)"`                // 昵称
}

// TableName 指定表名
func (Test100mCrc32Table) TableName() string {
	return "test_100m_crc32_table"
}
