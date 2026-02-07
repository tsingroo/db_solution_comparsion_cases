## 使用 PostgreSQL UUID 作为主键，插入一万条数据测试，每100条批量插入

### 配置
- 腾讯云
- PostgreSQL 版本：16.11
- 规格: S5.LARGE4
- 4C4G
- 100G云盘,3300IOPS,120MiB/s带宽

### 观察结果
- 连续执行多组CRUD任务，MySQL使用SnowflakeID作为主键在CRUD操作时会占用IOPS上限的一半，Postgresql使用UUID类型会将IOPS占满。
- 平时低负载的时候，MySQL使用snowflakeID执行CRUD和PostgreSQL使用UUID占用的IOPS都会在1600左右。
- 鉴于目前高负载时MySQL的磁盘IOPS还有空闲，可能MySQL使用snowflakeId在同样的IOPS下性能要好于PostgreSQL。只是在本测试程序没有将IOPS打满。



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
Create 完成，耗时: 1388 ms
Get 完成，耗时: 1135 ms
Update 完成，耗时: 1513 ms
Delete 完成，耗时: 1430 ms

Create 完成，耗时: 1581 ms
Get 完成，耗时: 1412 ms
Update 完成，耗时: 1658 ms
Delete 完成，耗时: 1661 ms

Create 完成，耗时: 1663 ms
Get 完成，耗时: 1251 ms
Update 完成，耗时: 1542 ms
Delete 完成，耗时: 1580 ms


### 亿级数据
#### 一万次只使用uuid作为主键，三轮测试
Create 完成，耗时: 1668 ms
Get 完成，耗时: 1406 ms
Update 完成，耗时: 1818 ms
Delete 完成，耗时: 1823 ms

Create 完成，耗时: 1802 ms
Get 完成，耗时: 1486 ms
Update 完成，耗时: 1885 ms
Delete 完成，耗时: 1629 ms

Create 完成，耗时: 1952 ms
Get 完成，耗时: 1226 ms
Update 完成，耗时: 2130 ms
Delete 完成，耗时: 1521 ms

