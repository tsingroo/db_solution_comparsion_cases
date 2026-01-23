## Db Optimization Techs - 一些数据库优化的技巧



### 场景1
- 一个表有四个字段(uuid, name, email, nickname)
- 对比使用单个uuid主键和使用【crc32(uuid)+uuid】联合主键时的在【空表、百万、亿】的【查询、插入、更新】性能差异
