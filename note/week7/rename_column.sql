-- Active: 1756683664904@@127.0.0.1@3306@zero_demo
-- 将user表中的id列重命名为user_id
USE zero_demo;

-- 修改列名，同时保持原有的数据类型和约束
ALTER TABLE `user`
    CHANGE COLUMN `age` `user_id` bigint NOT NULL AUTO_INCREMENT;

-- 注意：修改主键列名后，可能需要重新设置主键约束
-- 如果上述语句执行后主键约束丢失，可以执行以下语句重新设置
-- ALTER TABLE `user`
--    ADD PRIMARY KEY (`user_id`);