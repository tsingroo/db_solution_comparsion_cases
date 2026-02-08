## 使用 PostgreSQL snowflake Id 作为主键，插入一万条数据测试

### 配置
- 腾讯云
- PostgreSQL 版本：16.11
- 规格: S5.LARGE4
- 4C4G
- 100G云盘,3300IOPS,120MiB/s带宽

### 观察结果
- 都是使用snowflakeID作为主键的情况下，Postgresql在查询时要比MySQL慢20%-40%，但是在插入、删除、更新操作的时候要节约20%-40%的时间。
- 同样一万次的CRUD，Postgresql比MySQL占用的IOPS要少一些。Postgresql连续多次CRUD会占用1200左右的磁盘IOPS，MySQL会占到1600左右。



### 测试数据
### 零数据
#### 一万次只使用uuid作为主键，三轮测试
Create 完成，耗时: 1792 ms
Get 完成，耗时: 1327 ms
Update 完成，耗时: 1769 ms
Delete 完成，耗时: 1480 ms

Create 完成，耗时: 1594 ms
Get 完成，耗时: 1194 ms
Update 完成，耗时: 1499 ms
Delete 完成，耗时: 1466 ms

Create 完成，耗时: 1540 ms
Get 完成，耗时: 1274 ms
Update 完成，耗时: 1666 ms
Delete 完成，耗时: 1416 ms

### 百万级数据
#### 一万次只使用uuid作为主键，三轮测试
Create 完成，耗时: 1742 ms
Get 完成，耗时: 1301 ms
Update 完成，耗时: 1591 ms
Delete 完成，耗时: 1601 ms

Create 完成，耗时: 1445 ms
Get 完成，耗时: 1438 ms
Update 完成，耗时: 1426 ms
Delete 完成，耗时: 1313 ms

Create 完成，耗时: 1422 ms
Get 完成，耗时: 1264 ms
Update 完成，耗时: 1391 ms
Delete 完成，耗时: 1327 ms


### 亿级数据
#### 一万次只使用uuid作为主键，三轮测试
Create 完成，耗时: 1475 ms
Get 完成，耗时: 1287 ms
Update 完成，耗时: 1337 ms
Delete 完成，耗时: 1436 ms

Create 完成，耗时: 1576 ms
Get 完成，耗时: 1211 ms
Update 完成，耗时: 1501 ms
Delete 完成，耗时: 1675 ms

Create 完成，耗时: 1555 ms
Get 完成，耗时: 1175 ms
Update 完成，耗时: 1612 ms
Delete 完成，耗时: 1328 ms
