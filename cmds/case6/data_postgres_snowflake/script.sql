-- 创建数据库
CREATE DATABASE test_postgres_snowflake_db;

-- 连接到数据库
\c test_postgres_snowflake_db

-- 创建表（主键为 BIGINT，对应 Go 的 int64/Snowflake ID）
CREATE TABLE test_snowflake_table (
    id BIGINT PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(50),
    nickname VARCHAR(50)
);
