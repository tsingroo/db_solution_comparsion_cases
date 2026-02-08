package dals

import (
	"db_optimization_techs/pkgs/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TestPostgresUuidDAL 数据访问层，用于操作 test_postgres_uuid_table 表
type TestPostgresUuidDAL struct {
	db *gorm.DB
}

// NewTestPostgresUuidDAL 创建 TestPostgresUuidDAL 实例
func NewTestPostgresUuidDAL(db *gorm.DB) *TestPostgresUuidDAL {
	return &TestPostgresUuidDAL{db: db}
}

// Create 创建记录，若 record.ID 为空（uuid.Nil）则自动生成 UUID
func (dal *TestPostgresUuidDAL) Create(record *models.TestPostgresUuidTable) error {
	if record.ID == uuid.Nil {
		record.ID = uuid.New()
	}
	return dal.db.Create(record).Error
}

// GetByID 根据主键 id 查询记录
func (dal *TestPostgresUuidDAL) GetByID(id uuid.UUID) (*models.TestPostgresUuidTable, error) {
	var record models.TestPostgresUuidTable
	err := dal.db.Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// Update 按主键 id 定位并更新 name/email/nickname
func (dal *TestPostgresUuidDAL) Update(record *models.TestPostgresUuidTable) error {
	return dal.db.Model(&models.TestPostgresUuidTable{}).
		Where("id = ?", record.ID).
		Updates(map[string]interface{}{
			"name":     record.Name,
			"email":    record.Email,
			"nickname": record.Nickname,
		}).Error
}

// Delete 按主键 id 删除记录
func (dal *TestPostgresUuidDAL) Delete(id uuid.UUID) error {
	return dal.db.Where("id = ?", id).Delete(&models.TestPostgresUuidTable{}).Error
}

// InsertBatch100 批量插入多条记录（典型用法为 100 条），由调用方保证 records 长度与内容；DAL 内不强制校验 len(records)==100
// 使用 CreateInBatches 确保每批 100 行对应一条 INSERT，实现真正的批量插入
// 若记录的 ID 为空（uuid.Nil），则自动生成 UUID（保持与 Create 方法一致的行为）
func (dal *TestPostgresUuidDAL) InsertBatch100(records []*models.TestPostgresUuidTable) error {
	// 为所有 ID 为空的记录生成 UUID
	for _, record := range records {
		if record.ID == uuid.Nil {
			record.ID = uuid.New()
		}
	}
	// 批量插入
	return dal.db.CreateInBatches(records, 100).Error
}
