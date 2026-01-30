package dals

import (
	"db_optimization_techs/pkgs/models"

	"github.com/sony/sonyflake/v2"
	"gorm.io/gorm"
)

// TestSnowflakeDAL 数据访问层，用于操作 test_snowflake_table 表
type TestSnowflakeDAL struct {
	db *gorm.DB
	sf *sonyflake.Sonyflake
}

// NewTestSnowflakeDAL 创建 TestSnowflakeDAL 实例，内部创建 Sonyflake（MachineID 为 nil 时使用默认）
func NewTestSnowflakeDAL(db *gorm.DB) *TestSnowflakeDAL {
	sf, err := sonyflake.New(sonyflake.Settings{})
	if err != nil {
		panic("创建 Sonyflake 实例失败: " + err.Error())
	}
	return &TestSnowflakeDAL{db: db, sf: sf}
}

// Create 创建记录，若 record.ID == 0 则用 Sonyflake 生成 ID
func (dal *TestSnowflakeDAL) Create(record *models.TestSnowflakeTable) error {
	if record.ID == 0 {
		id, err := dal.sf.NextID()
		if err != nil {
			return err
		}
		record.ID = id
	}
	return dal.db.Create(record).Error
}

// GetByID 根据主键 id 查询记录
func (dal *TestSnowflakeDAL) GetByID(id int64) (*models.TestSnowflakeTable, error) {
	var record models.TestSnowflakeTable
	err := dal.db.Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

// Update 按主键 id 定位并更新 name/email/nickname
func (dal *TestSnowflakeDAL) Update(record *models.TestSnowflakeTable) error {
	return dal.db.Model(&models.TestSnowflakeTable{}).
		Where("id = ?", record.ID).
		Updates(map[string]interface{}{
			"name":     record.Name,
			"email":    record.Email,
			"nickname": record.Nickname,
		}).Error
}

// Delete 按主键 id 删除记录
func (dal *TestSnowflakeDAL) Delete(id int64) error {
	return dal.db.Where("id = ?", id).Delete(&models.TestSnowflakeTable{}).Error
}
