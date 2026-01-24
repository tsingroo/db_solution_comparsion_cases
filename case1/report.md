## 只使用uuid的主键和使用crc32+uuid的主键的CRUD的性能对比

### 配置
- 腾讯云
- 规格: S5.LARGE4
- 4C4G
- 100G云盘,3300IOPS,120MiB/s带宽



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
#### 一万次只使用uuid作为主键

#### 一万次使用crc32+uuid作为主键