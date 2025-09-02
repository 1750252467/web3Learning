-- 正确将age列重命名为user_id并设置为主键自增列
USE zero_demo;

-- 1. 先查看当前表结构，了解现有的自增列和主键
-- DESC `user`;

-- 2. 如果表中已有其他自增列，需要先移除其自增属性
-- 假设id列是当前的自增列
ALTER TABLE `user`
    MODIFY COLUMN `id` bigint NOT NULL;

-- 3. 将age列重命名为user_id，并设置为自增列和主键
ALTER TABLE `user`
    CHANGE COLUMN `age` `user_id` bigint NOT NULL AUTO_INCREMENT,
    DROP PRIMARY KEY,
    ADD PRIMARY KEY (`user_id`);

-- 4. 如果需要，可以删除原来的id列
-- ALTER TABLE `user`
--    DROP COLUMN `id`;

-- 5. 验证更改是否成功
-- DESC `user`;

-- 注意：这个脚本假设原来的id列是自增列。如果实际情况不同，
-- 请先查看表结构，根据实际情况调整脚本。