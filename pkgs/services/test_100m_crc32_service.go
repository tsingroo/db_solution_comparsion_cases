package services

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"sync"
	"time"

	"db_optimization_techs/pkgs/dals"
	"db_optimization_techs/pkgs/models"
	"github.com/google/uuid"
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

			id := uuid.New().String()
			record := &models.Test100mCrc32Table{
				Uuid:     id,
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

// Get 先创建 1 万条测试数据，然后随机查询 1 万次，返回总耗时（毫秒）
func (s *Test100mCrc32Service) Get() (int64, error) {
	// 准备阶段：创建 10000 条记录（不计时）
	uuids := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		id := uuid.New().String()
		record := &models.Test100mCrc32Table{
			Uuid:     id,
			Name:     fmt.Sprintf("TestName_%d", i),
			Email:    fmt.Sprintf("test_%d@test.com", i),
			Nickname: fmt.Sprintf("TestNickname_%d", i),
		}

		if err := s.dal.Create(record); err != nil {
			return 0, fmt.Errorf("创建测试数据失败: %w", err)
		}
		uuids = append(uuids, id)
	}

	// 随机打乱 UUID 列片
	rand.Shuffle(len(uuids), func(i, j int) {
		uuids[i], uuids[j] = uuids[j], uuids[i]
	})

	// 测试阶段：随机查询 10000 次（计时）
	start := time.Now()

	const maxConcurrency = 80
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	for _, uuid := range uuids {
		wg.Add(1)
		sem <- struct{}{}

		go func(u string) {
			defer wg.Done()
			defer func() { <-sem }()

			crc32Value := crc32.ChecksumIEEE([]byte(u))
			_, err := s.dal.GetByCrc32AndUUID(crc32Value, u)
			if err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(uuid)
	}

	wg.Wait()

	if len(errors) > 0 {
		return 0, fmt.Errorf("查询完成，但有 %d 个失败: %v", len(errors), errors[0])
	}

	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}

// Update 先创建 1 万条测试数据，然后循环更新 1 万次，返回总耗时（毫秒）
func (s *Test100mCrc32Service) Update() (int64, error) {
	// 准备阶段：创建 10000 条记录（不计时）
	uuids := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		id := uuid.New().String()
		record := &models.Test100mCrc32Table{
			Uuid:     id,
			Name:     fmt.Sprintf("OriginalName_%d", i),
			Email:    fmt.Sprintf("original_%d@test.com", i),
			Nickname: fmt.Sprintf("OriginalNickname_%d", i),
		}

		if err := s.dal.Create(record); err != nil {
			return 0, fmt.Errorf("创建测试数据失败: %w", err)
		}
		uuids = append(uuids, id)
	}

	// 测试阶段：循环更新 10000 次（计时）
	start := time.Now()

	const maxConcurrency = 80
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	for i, uuid := range uuids {
		wg.Add(1)
		sem <- struct{}{}

		go func(index int, u string) {
			defer wg.Done()
			defer func() { <-sem }()

			updateRecord := &models.Test100mCrc32Table{
				Uuid:     u,
				Name:     fmt.Sprintf("UpdatedName_%d", index),
				Email:    fmt.Sprintf("updated_%d@test.com", index),
				Nickname: fmt.Sprintf("UpdatedNickname_%d", index),
			}

			if err := s.dal.Update(updateRecord); err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(i, uuid)
	}

	wg.Wait()

	if len(errors) > 0 {
		return 0, fmt.Errorf("更新完成，但有 %d 个失败: %v", len(errors), errors[0])
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
		// 使用标准 UUID v4 生成唯一标识
		id := uuid.New().String()
		record := &models.Test100mCrc32Table{
			Uuid:     id,
			Name:     fmt.Sprintf("DeleteName_%d", i),
			Email:    fmt.Sprintf("delete_%d@test.com", i),
			Nickname: fmt.Sprintf("DeleteNickname_%d", i),
		}

		// 创建记录（不计时）
		if err := s.dal.Create(record); err != nil {
			return 0, fmt.Errorf("第 %d 次创建记录失败: %w", i+1, err)
		}
		uuids = append(uuids, id)
	}

	// 删除阶段：删除所有记录（只统计这部分时间）
	start := time.Now()

	const maxConcurrency = 80
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	for _, uuid := range uuids {
		wg.Add(1)
		sem <- struct{}{}

		go func(u string) {
			defer wg.Done()
			defer func() { <-sem }()

			if err := s.dal.Delete(u); err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(uuid)
	}

	wg.Wait()

	if len(errors) > 0 {
		return 0, fmt.Errorf("删除完成，但有 %d 个失败: %v", len(errors), errors[0])
	}

	elapsed := time.Since(start)
	return elapsed.Milliseconds(), nil
}
