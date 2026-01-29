package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"db_optimization_techs/pkgs/dals"
	"db_optimization_techs/pkgs/models"

	"github.com/google/uuid"
)

// Test100mService 服务层，用于测试 Test100mDAL 的性能
type Test100mService struct {
	dal *dals.Test100mDAL
}

// NewTest100mService 创建 Test100mService 实例
func NewTest100mService(dal *dals.Test100mDAL) *Test100mService {
	return &Test100mService{dal: dal}
}

// InsertBatch10000 批量插入 10000 条：并行 100 批，每批在 Service 内生成 100 条并调用 DAL.InsertBatch100，返回总耗时（毫秒）
func (s *Test100mService) InsertBatch10000() (int64, error) {
	start := time.Now()
	const batchSize = 100
	const loopCount = 100
	const maxConcurrency = 30 // 有界并发，避免打满 DB 连接池
	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstErr error
	for batch := 0; batch < loopCount; batch++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(batch int) {
			defer wg.Done()
			defer func() { <-sem }()
			records := make([]*models.Test100mTable, 0, batchSize)
			for i := 0; i < batchSize; i++ {
				globalIdx := batch*batchSize + i
				records = append(records, &models.Test100mTable{
					Uuid:     uuid.New().String(),
					Name:     fmt.Sprintf("Name_%d", globalIdx),
					Email:    fmt.Sprintf("email_%d@test.com", globalIdx),
					Nickname: fmt.Sprintf("Nickname_%d", globalIdx),
				})
			}
			if err := s.dal.InsertBatch100(records); err != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = err
				}
				mu.Unlock()
			}
		}(batch)
	}
	wg.Wait()
	if firstErr != nil {
		return time.Since(start).Milliseconds(), firstErr
	}
	return time.Since(start).Milliseconds(), nil
}

// Create 循环 1 万次创建记录，返回总耗时（毫秒）
func (s *Test100mService) Create() (int64, error) {
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
			record := &models.Test100mTable{
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
func (s *Test100mService) Get() (int64, error) {
	// 准备阶段：创建 10000 条记录（不计时）
	uuids := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		id := uuid.New().String()
		record := &models.Test100mTable{
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

			_, err := s.dal.GetByUUID(u)
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
func (s *Test100mService) Update() (int64, error) {
	// 准备阶段：创建 10000 条记录（不计时）
	uuids := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		id := uuid.New().String()
		record := &models.Test100mTable{
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

			updateRecord := &models.Test100mTable{
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
func (s *Test100mService) Delete() (int64, error) {
	// 准备阶段：创建 10000 条记录（不计时）
	uuids := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		// 使用标准 UUID v4 生成唯一标识
		id := uuid.New().String()
		record := &models.Test100mTable{
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
