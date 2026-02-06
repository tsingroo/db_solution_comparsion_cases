## 使用snowflake生成的ID作为主键

### 配置
- 腾讯云
- 规格: S5.LARGE4
- 4C4G
- 100G云盘,3300IOPS,120MiB/s带宽

### 观察结果
- 虽然在数据量比较少，使用snowflake生成ID在当前程序下插入性能差不多，但是IOPS只有1600只使用了一半左右，还有很大增长空间，IOPS在当前程序下不再是瓶颈



### 零数据
#### 一万次只使用uuid作为主键，三轮测试
Create 完成，耗时: 2007 ms
Get 完成，耗时: 772 ms
Update 完成，耗时: 2049 ms
Delete 完成，耗时: 1984 ms

Create 完成，耗时: 1988 ms
Get 完成，耗时: 769 ms
Update 完成，耗时: 2032 ms
Delete 完成，耗时: 2005 ms

Create 完成，耗时: 1894 ms
Get 完成，耗时: 801 ms
Update 完成，耗时: 2106 ms
Delete 完成，耗时: 1948 ms

### 百万级数据
#### 一万次只使用uuid作为主键，三轮测试
Create 完成，耗时: 1882 ms
Get 完成，耗时: 768 ms
Update 完成，耗时: 2055 ms
Delete 完成，耗时: 1981 ms

Create 完成，耗时: 1793 ms
Get 完成，耗时: 802 ms
Update 完成，耗时: 2046 ms
Delete 完成，耗时: 1948 ms

Create 完成，耗时: 1748 ms
Get 完成，耗时: 779 ms
Update 完成，耗时: 1978 ms
Delete 完成，耗时: 1935 m

### 亿级数据
#### 一万次只使用uuid作为主键，三轮测试