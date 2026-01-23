package services

import (
	"fmt"
	"time"

	"db_optimization_techs/pkgs/dals"
	"db_optimization_techs/pkgs/models"
)

// Test100mService 服务层，用于测试 Test100mDAL 的性能
type Test100mService struct {
	dal *dals.Test100mDAL
}

// NewTest100mService 创建 Test100mService 实例
func NewTest100mService(dal *dals.Test100mDAL) *Test100mService {
	return &Test100mService{dal: dal}
}

// Create 循环 1 万次创建记录，返回总耗时（毫秒）
func (s *Test100mService) Create() (int64, error) {
	start := time.Now()

	for i := 0; i < 10000; i++ {
		// 生成唯一的 UUID 和测试数据
		uuid := fmt.Sprintf("test-uuid-%d-%d", time.Now().UnixNano(), i)
		record := &models.Test100mTable{
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

// Get 先创建 1 条测试数据，然后循环 1 万次查询，返回总耗时（毫秒）
func (s *Test100mService) Get() (int64, error) {
	// 先创建测试数据
	testUUID := fmt.Sprintf("test-get-uuid-%d", time.Now().UnixNano())
	testRecord := &models.Test100mTable{
		Uuid:     testUUID,
		Name:     "TestName",
		Email:    "test@test.com",
		Nickname: "TestNickname",
	}

	if err := s.dal.Create(testRecord); err != nil {
		return 0, fmt.Errorf("创建测试数据失败: %w", err)
	}

	// 开始计时并循环查询
	start := time.Now()

	for i := 0; i < 10000; i++ {
		_, err := s.dal.GetByUUID(testUUID)
		if err != nil {
			return 0, fmt.Errorf("第 %d 次查询失败: %w", i+1, err)
		}
	}

	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}

// Update 先创建 1 条测试数据，然后循环 1 万次更新，返回总耗时（毫秒）
func (s *Test100mService) Update() (int64, error) {
	// 先创建测试数据
	testUUID := fmt.Sprintf("test-update-uuid-%d", time.Now().UnixNano())
	testRecord := &models.Test100mTable{
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
		updateRecord := &models.Test100mTable{
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
func (s *Test100mService) Delete() (int64, error) {
	// 准备阶段：创建 10000 条记录（不计时）
	uuids := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		// 生成唯一的 UUID 和测试数据
		uuid := fmt.Sprintf("test-delete-uuid-%d-%d", time.Now().UnixNano(), i)
		record := &models.Test100mTable{
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
