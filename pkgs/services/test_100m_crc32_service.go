package services

import (
	"fmt"
	"hash/crc32"
	"time"

	"db_optimization_techs/pkgs/dals"
	"db_optimization_techs/pkgs/models"
)

// Test100mCrc32Service 服务层，用于测试 Test100mCrc32DAL 的性能
type Test100mCrc32Service struct {
	dal *dals.Test100mCrc32DAL
}

// NewTest100mCrc32Service 创建 Test100mCrc32Service 实例
func NewTest100mCrc32Service(dal *dals.Test100mCrc32DAL) *Test100mCrc32Service {
	return &Test100mCrc32Service{dal: dal}
}

// Create 循环 1 万次创建记录，返回总耗时（毫秒）
// CRC32 值会在 DAL 层自动计算
func (s *Test100mCrc32Service) Create() (int64, error) {
	start := time.Now()
	
	for i := 0; i < 10000; i++ {
		// 生成唯一的 UUID 和测试数据
		uuid := fmt.Sprintf("test-uuid-%d-%d", time.Now().UnixNano(), i)
		record := &models.Test100mCrc32Table{
			Uuid:     uuid,
			Name:     fmt.Sprintf("Name_%d", i),
			Email:    fmt.Sprintf("email_%d@test.com", i),
			Nickname: fmt.Sprintf("Nickname_%d", i),
		}
		
		if err := s.dal.Create(record); err != nil {
			return 0, fmt.Errorf("第 %d 次创建失败: %w", i+1, err)
		}
	}
	
	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}

// Get 先创建 1 条测试数据并计算 CRC32，然后循环 1 万次查询，返回总耗时（毫秒）
func (s *Test100mCrc32Service) Get() (int64, error) {
	// 先创建测试数据
	testUUID := fmt.Sprintf("test-get-uuid-%d", time.Now().UnixNano())
	testRecord := &models.Test100mCrc32Table{
		Uuid:     testUUID,
		Name:     "TestName",
		Email:    "test@test.com",
		Nickname: "TestNickname",
	}
	
	if err := s.dal.Create(testRecord); err != nil {
		return 0, fmt.Errorf("创建测试数据失败: %w", err)
	}
	
	// 计算 CRC32 值
	crc32Value := crc32.ChecksumIEEE([]byte(testUUID))
	
	// 开始计时并循环查询
	start := time.Now()
	
	for i := 0; i < 10000; i++ {
		_, err := s.dal.GetByCrc32AndUUID(crc32Value, testUUID)
		if err != nil {
			return 0, fmt.Errorf("第 %d 次查询失败: %w", i+1, err)
		}
	}
	
	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}

// Update 先创建 1 条测试数据，然后循环 1 万次更新，返回总耗时（毫秒）
func (s *Test100mCrc32Service) Update() (int64, error) {
	// 先创建测试数据
	testUUID := fmt.Sprintf("test-update-uuid-%d", time.Now().UnixNano())
	testRecord := &models.Test100mCrc32Table{
		Uuid:     testUUID,
		Name:     "OriginalName",
		Email:    "original@test.com",
		Nickname: "OriginalNickname",
	}
	
	if err := s.dal.Create(testRecord); err != nil {
		return 0, fmt.Errorf("创建测试数据失败: %w", err)
	}
	
	// 开始计时并循环更新
	start := time.Now()
	
	for i := 0; i < 10000; i++ {
		updateRecord := &models.Test100mCrc32Table{
			Uuid:     testUUID,
			Name:     fmt.Sprintf("UpdatedName_%d", i),
			Email:    fmt.Sprintf("updated_%d@test.com", i),
			Nickname: fmt.Sprintf("UpdatedNickname_%d", i),
		}
		
		if err := s.dal.Update(updateRecord); err != nil {
			return 0, fmt.Errorf("第 %d 次更新失败: %w", i+1, err)
		}
	}
	
	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}

// Delete 先创建 1 万条记录，然后删除这 1 万条记录，返回总耗时（毫秒）
// 只统计删除操作的时间，不包含创建记录的时间
func (s *Test100mCrc32Service) Delete() (int64, error) {
	// 准备阶段：创建 10000 条记录（不计时）
	uuids := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		// 生成唯一的 UUID 和测试数据
		uuid := fmt.Sprintf("test-delete-uuid-%d-%d", time.Now().UnixNano(), i)
		record := &models.Test100mCrc32Table{
			Uuid:     uuid,
			Name:     fmt.Sprintf("DeleteName_%d", i),
			Email:    fmt.Sprintf("delete_%d@test.com", i),
			Nickname: fmt.Sprintf("DeleteNickname_%d", i),
		}
		
		// 创建记录（不计时）
		if err := s.dal.Create(record); err != nil {
			return 0, fmt.Errorf("第 %d 次创建记录失败: %w", i+1, err)
		}
		uuids = append(uuids, uuid)
	}
	
	// 删除阶段：删除所有记录（只统计这部分时间）
	start := time.Now()
	for i, uuid := range uuids {
		if err := s.dal.Delete(uuid); err != nil {
			return 0, fmt.Errorf("第 %d 次删除失败: %w", i+1, err)
		}
	}
	elapsed := time.Since(start)
	
	return elapsed.Milliseconds(), nil
}
