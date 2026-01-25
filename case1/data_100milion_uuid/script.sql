-- 创建数据库 test_100m_db
CREATE DATABASE IF NOT EXISTS test_100m_db
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE test_100m_db;

-- 创建表 test_100m_table
CREATE TABLE IF NOT EXISTS test_100m_table (
    uuid VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(50),
    nickname VARCHAR(50)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci;
