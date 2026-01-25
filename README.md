## Some Db cases - 一些数据库的不同场景下索引选择的实际对比



### 场景1
- 数据库使用MySQL 8.0.44
- 一个表有四个字段(uuid, name, email, nickname)
- 对比使用单个uuid主键和使用【crc32(uuid)+uuid】联合主键时的在【空表、百万、亿】的【查询、插入、更新】性能差异
