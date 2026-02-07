## 使用 PostgreSQL UUID 作为主键，插入一万条数据测试，每100条批量插入

### 配置
- 腾讯云
- PostgreSQL 版本：16.11
- 规格: S5.LARGE4
- 4C4G
- 100G云盘,3300IOPS,120MiB/s带宽

### 观察结果
- [测试后填写观察结果]

### 测试数据
### 零数据
#### 一万次只使用uuid作为主键，三轮测试
Create 完成，耗时: 1371 ms
Get 完成，耗时: 1489 ms
Update 完成，耗时: 1644 ms
Delete 完成，耗时: 1386 ms

Create 完成，耗时: 1402 ms
Get 完成，耗时: 1313 ms
Update 完成，耗时: 1526 ms
Delete 完成，耗时: 1579 ms

Create 完成，耗时: 1618 ms
Get 完成，耗时: 1264 ms
Update 完成，耗时: 1688 ms
Delete 完成，耗时: 1366 ms


### 百万级数据
#### 一万次只使用uuid作为主键，三轮测试


### 亿级数据
#### 一万次只使用uuid作为主键，三轮测试


