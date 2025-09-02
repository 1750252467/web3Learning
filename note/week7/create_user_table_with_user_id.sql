-- 新建user表并增加user_id和nickname字段
-- 方案1：保留id作为自增主键，添加普通的user_id列
USE zero_demo;

CREATE TABLE user(
    id int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'Primary Key',
    create_time DATETIME COMMENT 'Create Time',
    name VARCHAR(255),
    user_id VARCHAR(50) NOT NULL COMMENT '用户唯一标识',
    nickname VARCHAR(100) COMMENT '用户昵称'
) COMMENT 'user';

-- 方案2：将user_id设为自增主键（如果需要以user_id为主键）
-- 请根据实际需求选择其中一个方案执行
-- USE zero_demo;
-- 
-- CREATE TABLE user(
--     id int NOT NULL COMMENT '原始ID',
--     create_time DATETIME COMMENT 'Create Time',
--     name VARCHAR(255),
--     user_id bigint NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '用户唯一标识（主键）',
--     nickname VARCHAR(100) COMMENT '用户昵称'
-- ) COMMENT 'user';

-- 注意：请不要同时执行两个方案，只选择其中一个适合您需求的方案。