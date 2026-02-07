-- 创建数据库
CREATE DATABASE test_postgres_uuid_db;

-- 连接到数据库
\c test_postgres_uuid_db

-- 创建表（使用 PostgreSQL 内置 UUID 类型）
-- 注意：UUID 由应用层生成，不使用 DEFAULT 约束
CREATE TABLE test_postgres_uuid_table (
    id UUID PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(50),
    nickname VARCHAR(50)
);
