-- Active: 1756683664904@@127.0.0.1@3306@zero_demo
-- 向user表中添加user_id列
USE zero_demo;

-- 添加user_id列，设置为bigint类型，非空且自增
-- 注意：如果表中已有数据，需要处理如何填充现有行的user_id值
ALTER TABLE `user`
    ADD COLUMN `user_id` bigint NOT NULL AUTO_INCREMENT FIRST,
    DROP PRIMARY KEY,
    ADD PRIMARY KEY (`user_id`);

-- 如果需要保留原有的id列，可以不执行下面的语句
-- 如果确实需要移除原有的id列，可以取消下面的注释
-- ALTER TABLE `user`
--    DROP COLUMN `id`;