package dals

import (
	"hash/crc32"

	"db_optimization_techs/pkgs/models"

	"gorm.io/gorm"
)

// Test100mCrc32DAL 数据访问层，用于操作 test_100m_crc32_table 表
type Test100mCrc32DAL struct {
	db *gorm.DB
}

// NewTest100mCrc32DAL 创建 Test100mCrc32DAL 实例
func NewTest100mCrc32DAL(db *gorm.DB) *Test100mCrc32DAL {
	return &Test100mCrc32DAL{db: db}
}

// Create 创建记录，自动计算 uuid_crc32
func (dal *Test100mCrc32DAL) Create(record *models.Test100mCrc32Table) error {
	// 自动计算 UUID 的 CRC32 值
	record.UuidCrc32 = crc32.ChecksumIEEE([]byte(record.Uuid))
	return dal.db.Create(record).Error
}

// GetByCrc32AndUUID 根据 CRC32 和 UUID 查询记录（直接使用联合主键）
func (dal *Test100mCrc32DAL) GetByCrc32AndUUID(crc32 uint32, uuid string) (*models.Test100mCrc32Table, error) {
	var record models.Test100mCrc32Table
	err := dal.db.Where("uuid_crc32 = ? AND uuid = ?", crc32, uuid).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// Update 更新记录，自动更新 uuid_crc32（使用联合主键定位）
func (dal *Test100mCrc32DAL) Update(record *models.Test100mCrc32Table) error {
	// 如果 UUID 发生变化，重新计算 CRC32
	record.UuidCrc32 = crc32.ChecksumIEEE([]byte(record.Uuid))
	// 使用联合主键 (uuid_crc32, uuid) 定位记录并更新
	return dal.db.Model(&models.Test100mCrc32Table{}).
		Where("uuid_crc32 = ? AND uuid = ?", record.UuidCrc32, record.Uuid).
		Updates(map[string]interface{}{
			"name":     record.Name,
			"email":    record.Email,
			"nickname": record.Nickname,
		}).Error
}

// Delete 删除记录（使用联合主键 (uuid_crc32, uuid) 定位）
func (dal *Test100mCrc32DAL) Delete(uuid string) error {
	// 计算 CRC32 后使用联合主键删除
	crc32Value := crc32.ChecksumIEEE([]byte(uuid))
	// 明确使用联合主键索引进行删除
	return dal.db.Model(&models.Test100mCrc32Table{}).
		Where("uuid_crc32 = ? AND uuid = ?", crc32Value, uuid).
		Delete(&models.Test100mCrc32Table{}).Error
}
