-- 创建数据库 test_snowflake_db
CREATE DATABASE IF NOT EXISTS test_snowflake_db
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE test_snowflake_db;

-- 创建表 test_snowflake_table（主键为 Snowflake 算法生成的 int64）
CREATE TABLE IF NOT EXISTS test_snowflake_table (
    id BIGINT,
    name VARCHAR(50),
    email VARCHAR(50),
    nickname VARCHAR(50),
    PRIMARY KEY (id)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci;
