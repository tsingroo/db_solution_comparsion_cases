package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"db_optimization_techs/pkgs/dals"
	"db_optimization_techs/pkgs/models"
)

// TestSnowflakeService 服务层，用于测试 TestSnowflakeDAL 的性能
type TestSnowflakeService struct {
	dal *dals.TestSnowflakeDAL
}

// NewTestSnowflakeService 创建 TestSnowflakeService 实例
func NewTestSnowflakeService(dal *dals.TestSnowflakeDAL) *TestSnowflakeService {
	return &TestSnowflakeService{dal: dal}
}

// Create 并发 1 万次创建记录，ID 由 DAL 内 Sonyflake 生成，返回总耗时（毫秒）
func (s *TestSnowflakeService) Create() (int64, error) {
	start := time.Now()

	const maxConcurrency = 80
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(index int) {
			defer wg.Done()
			defer func() { <-sem }()

			record := &models.TestSnowflakeTable{
				Name:     fmt.Sprintf("Name_%d", index),
				Email:    fmt.Sprintf("email_%d@test.com", index),
				Nickname: fmt.Sprintf("Nickname_%d", index),
			}

			if err := s.dal.Create(record); err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()

	if len(errors) > 0 {
		return 0, fmt.Errorf("创建完成，但有 %d 个失败: %v", len(errors), errors[0])
	}

	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}

// Get 先创建 1 万条测试数据并收集 ID，再随机按 ID 查询 1 万次，返回总耗时（毫秒）
func (s *TestSnowflakeService) Get() (int64, error) {
	// 准备阶段：创建 10000 条记录（不计时），每次 Create 后从 record.ID 收集 ID
	ids := make([]int64, 0, 10000)
	for i := 0; i < 10000; i++ {
		record := &models.TestSnowflakeTable{
			Name:     fmt.Sprintf("TestName_%d", i),
			Email:    fmt.Sprintf("test_%d@test.com", i),
			Nickname: fmt.Sprintf("TestNickname_%d", i),
		}

		if err := s.dal.Create(record); err != nil {
			return 0, fmt.Errorf("创建测试数据失败: %w", err)
		}
		ids = append(ids, record.ID)
	}

	// 随机打乱 ID 切片
	rand.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})

	// 测试阶段：按 ID 查询 10000 次（计时）
	start := time.Now()

	const maxConcurrency = 80
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	for _, id := range ids {
		wg.Add(1)
		sem <- struct{}{}

		go func(rid int64) {
			defer wg.Done()
			defer func() { <-sem }()

			_, err := s.dal.GetByID(rid)
			if err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(id)
	}

	wg.Wait()

	if len(errors) > 0 {
		return 0, fmt.Errorf("查询完成，但有 %d 个失败: %v", len(errors), errors[0])
	}

	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}

// Update 先创建 1 万条测试数据并收集 ID，再按 ID 更新 1 万次，返回总耗时（毫秒）
func (s *TestSnowflakeService) Update() (int64, error) {
	// 准备阶段：创建 10000 条记录（不计时），收集 ID
	ids := make([]int64, 0, 10000)
	for i := 0; i < 10000; i++ {
		record := &models.TestSnowflakeTable{
			Name:     fmt.Sprintf("OriginalName_%d", i),
			Email:    fmt.Sprintf("original_%d@test.com", i),
			Nickname: fmt.Sprintf("OriginalNickname_%d", i),
		}

		if err := s.dal.Create(record); err != nil {
			return 0, fmt.Errorf("创建测试数据失败: %w", err)
		}
		ids = append(ids, record.ID)
	}

	// 测试阶段：按 ID 更新 10000 次（计时）
	start := time.Now()

	const maxConcurrency = 80
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	for i, id := range ids {
		wg.Add(1)
		sem <- struct{}{}

		go func(index int, rid int64) {
			defer wg.Done()
			defer func() { <-sem }()

			updateRecord := &models.TestSnowflakeTable{
				ID:       rid,
				Name:     fmt.Sprintf("UpdatedName_%d", index),
				Email:    fmt.Sprintf("updated_%d@test.com", index),
				Nickname: fmt.Sprintf("UpdatedNickname_%d", index),
			}

			if err := s.dal.Update(updateRecord); err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(i, id)
	}

	wg.Wait()

	if len(errors) > 0 {
		return 0, fmt.Errorf("更新完成，但有 %d 个失败: %v", len(errors), errors[0])
	}

	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}

// Delete 先创建 1 万条记录并收集 ID，再按 ID 删除这 1 万条，返回总耗时（毫秒）
// 只统计删除操作的时间，不包含创建记录的时间
func (s *TestSnowflakeService) Delete() (int64, error) {
	// 准备阶段：创建 10000 条记录（不计时），收集 ID
	ids := make([]int64, 0, 10000)
	for i := 0; i < 10000; i++ {
		record := &models.TestSnowflakeTable{
			Name:     fmt.Sprintf("DeleteName_%d", i),
			Email:    fmt.Sprintf("delete_%d@test.com", i),
			Nickname: fmt.Sprintf("DeleteNickname_%d", i),
		}

		if err := s.dal.Create(record); err != nil {
			return 0, fmt.Errorf("第 %d 次创建记录失败: %w", i+1, err)
		}
		ids = append(ids, record.ID)
	}

	// 删除阶段：按 ID 删除所有记录（只统计这部分时间）
	start := time.Now()

	const maxConcurrency = 80
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	for _, id := range ids {
		wg.Add(1)
		sem <- struct{}{}

		go func(rid int64) {
			defer wg.Done()
			defer func() { <-sem }()

			if err := s.dal.Delete(rid); err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(id)
	}

	wg.Wait()

	if len(errors) > 0 {
		return 0, fmt.Errorf("删除完成，但有 %d 个失败: %v", len(errors), errors[0])
	}

	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}
