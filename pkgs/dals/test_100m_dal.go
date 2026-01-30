package dals

import (
	"db_optimization_techs/pkgs/models"

	"gorm.io/gorm"
)

// Test100mDAL 数据访问层，用于操作 test_100m_table 表
type Test100mDAL struct {
	db *gorm.DB
}

// NewTest100mDAL 创建 Test100mDAL 实例
func NewTest100mDAL(db *gorm.DB) *Test100mDAL {
	return &Test100mDAL{db: db}
}

// Create 创建记录
func (dal *Test100mDAL) Create(record *models.Test100mTable) error {
	return dal.db.Create(record).Error
}

// InsertBatch100 批量插入多条记录（典型用法为 100 条），由调用方保证 records 长度与内容；DAL 内不强制校验 len(records)==100
// 使用 CreateInBatches 确保每批 100 行对应一条 INSERT，实现真正的批量插入
func (dal *Test100mDAL) InsertBatch100(records []*models.Test100mTable) error {
	return dal.db.CreateInBatches(records, 100).Error
}

// GetByUUID 根据 UUID 主键查询记录
func (dal *Test100mDAL) GetByUUID(uuid string) (*models.Test100mTable, error) {
	var record models.Test100mTable
	err := dal.db.Where("uuid = ?", uuid).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// Update 更新记录
func (dal *Test100mDAL) Update(record *models.Test100mTable) error {
	return dal.db.Save(record).Error
}

// Delete 根据 UUID 删除记录
func (dal *Test100mDAL) Delete(uuid string) error {
	return dal.db.Where("uuid = ?", uuid).Delete(&models.Test100mTable{}).Error
}
