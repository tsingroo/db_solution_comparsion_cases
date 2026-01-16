### 场景描述
- 一个表有四个字段(uuid, name, email, nickname)
- 数据量单表一亿条数据
- 对比使用单个uuid主键和使用【crc32(uuid)+uuid】联合主键时的查询、插入、更新性能差异


