## Some Db cases - 一些数据库的不同场景下索引选择的实际对比



### 场景1
- 数据库使用MySQL 8.0.44
- 一个表有四个字段(uuid, name, email, nickname)
- 对比使用单个uuid主键和使用【crc32(uuid)+uuid】联合主键时的在【空表、百万、亿】的【查询、插入、更新】性能差异有多大


### 场景2
- 数据库使用MySQL 8.0.44
- 表有四个字段(ID, name, email, nickname)，这次的ID使用snowflake算法生成
- 看看在【空表、百万、亿】的【查询、插入、更新】操作下，跟场景1的使用uuid作为主键的性能差异有多大



