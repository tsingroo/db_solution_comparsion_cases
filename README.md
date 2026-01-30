## Some Db cases - 一些数据库的不同场景下索引选择的实际对比



### 场景1
- 数据库使用MySQL 8.0.44
- 一个表有四个字段(uuid, name, email, nickname)
- 对比使用单个uuid主键和使用【crc32(uuid)+uuid】联合主键时的在【空表、百万、亿】的【查询、插入、更新】性能差异有多大


### 场景2
- 数据库使用MySQL 8.0.44
- 表有四个字段(uuid, name, email, nickname)
- 看看在【亿】级的【每100条批量插入】操作下，跟场景1的使用uuid单条插入的性能差异有多大


### 场景3
- 数据库使用MySQL 8.0.44
- 表有四个字段(ID(int64), name, email, nickname)，这次的ID使用snowflake算法生成
- 看看在【空表、百万、亿】的【查询、插入、更新】操作下，跟场景1的使用uuid作为主键的性能差异有多大


### 场景4
- 数据库使用MySQL 8.0.44
- 表有四个字段(ID(int64), name, email, nickname)，这次的ID使用snowflake算法生成
- 看看在【亿】级的【每100条批量插入】操作下，跟场景3的使用单条插入的性能差异有多大

### 场景5
- 数据库使用Postgresql
- 表有四个字段(ID(uuid类型), name, email, nickname)，




