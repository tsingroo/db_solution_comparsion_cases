## 只使用uuid的主键和使用crc32+uuid的主键的CRUD的性能对比

### 配置
- 腾讯云
- 规格: S5.LARGE4
- 4C4G
- 100G云盘,3300IOPS,120MiB/s带宽

### 观察结果
- 使用针对uuid添加crc32前缀的联合索引在各个量级上的CRUD速度都要比单一使用uuid作为主键的表要慢
- 磁盘IO是瓶颈插入，持续插入或更新会轻易将IO打满。CPU占用再20%左右,内存占用400M左右

#### 疑惑点
- 之前以为联合索引先使用crc32这个uint进行索引后再使用uuid索引，理论上应该会比只使用36位字符串的花费更小的代价，但是实际测试似乎不支持这个观点





### 零数据
#### 一万次只使用uuid作为主键
Create 完成，耗时: 1814 ms
Get 完成，耗时: 885 ms
Update 完成，耗时: 2007 ms
Delete 完成，耗时: 1850 ms

#### 一万次使用crc32+uuid作为主键
Create 完成，耗时: 1826 ms
Get 完成，耗时: 915 ms
Update 完成，耗时: 2064 ms
Delete 完成，耗时: 1949 ms

### 百万级数据
#### 一万次只使用uuid作为主键
Create 完成，耗时: 1858 ms
Get 完成，耗时: 864 ms
Update 完成，耗时: 2098 ms
Delete 完成，耗时: 1989 ms

#### 一万次使用crc32+uuid作为主键
Create 完成，耗时: 2001 ms
Get 完成，耗时: 908 ms
Update 完成，耗时: 2194 ms
Delete 完成，耗时: 2027 ms

### 亿级数据
#### 一万次只使用uuid作为主键，三轮测试
Create 完成，耗时: 2714 ms
Get 完成，耗时: 1209 ms
Update 完成，耗时: 4003 ms
Delete 完成，耗时: 3406 ms


Create 完成，耗时: 2841 ms
Get 完成，耗时: 1216 ms
Update 完成，耗时: 3924 ms
Delete 完成，耗时: 3210 ms


Create 完成，耗时: 2780 ms
Get 完成，耗时: 1187 ms
Update 完成，耗时: 3714 ms
Delete 完成，耗时: 3153 ms


#### 一万次使用crc32+uuid作为主键,三轮测试
Create 完成，耗时: 3216 ms
Get 完成，耗时: 1320 ms
Update 完成，耗时: 4402 ms
Delete 完成，耗时: 3654 ms


Create 完成，耗时: 3327 ms
Get 完成，耗时: 1449 ms
Update 完成，耗时: 4516 ms
Delete 完成，耗时: 3595 ms


Create 完成，耗时: 3586 ms
Get 完成，耗时: 1342 ms
Update 完成，耗时: 4449 ms
Delete 完成，耗时: 3596 ms