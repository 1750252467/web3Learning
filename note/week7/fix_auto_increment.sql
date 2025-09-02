-- 修复user表中的auto_increment和主键约束问题
USE zero_demo;

-- 查看当前表结构，了解现有列和约束情况
-- DESC `user`;

-- 解决多个auto_increment列的问题
-- 1. 先移除所有auto_increment属性
ALTER TABLE `user`
    MODIFY COLUMN `id` bigint NOT NULL,
    MODIFY COLUMN `user_id` bigint NOT NULL;

-- 2. 然后选择一个列作为auto_increment主键
-- 这里我们选择user_id作为auto_increment主键
ALTER TABLE `user`
    MODIFY COLUMN `user_id` bigint NOT NULL AUTO_INCREMENT,
    DROP PRIMARY KEY,
    ADD PRIMARY KEY (`user_id`);

-- 如果需要重命名列，可以取消下面的注释
-- ALTER TABLE `user`
--    CHANGE COLUMN `id` `old_id` bigint NOT NULL;

-- 如果确定不再需要原有的id列，可以取消下面的注释
-- ALTER TABLE `user`
--    DROP COLUMN `id`;

-- 验证更改是否成功
-- DESC `user`;