
#基本CRUD操作
#向 students 表中插入一条姓名为 "张三"，年龄为 20，年级为 "三年级" 的记录：
INSERT INTO students (name, age, grade) 
VALUES ('张三', 20, '三年级');



#编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
SELECT * FROM students 
WHERE age > 18;

#编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
UPDATE students 
SET grade = '四年级' 
WHERE name = '张三';

#删除 students 表中年龄小于 15 岁的学生记录：

DELETE FROM students 
WHERE age < 15;

#题目 2：事务语句 A：1 B:2  A余额：


BEGIN TRANSACTION;

-- 声明变量存储账户A的当前余额
DECLARE @A_Balance DECIMAL(10,2);

-- 获取账户A的当前余额
SELECT @A_Balance = balance FROM accounts WHERE id = [1];

-- 检查余额是否充足
IF @A_Balance >= 100
BEGIN
    -- 从账户A扣除100元
    UPDATE accounts 
    SET balance = balance - 100 
    WHERE id = [1];
    
    -- 向账户B增加100元
    UPDATE accounts 
    SET balance = balance + 100 
    WHERE id = [2];
    
    -- 记录转账信息
    INSERT INTO transactions (from_account_id, to_account_id, amount)
    VALUES ([1], [2], 100);
    
    -- 提交事务
    COMMIT TRANSACTION;
END
ELSE
BEGIN
    -- 余额不足，回滚事务
    ROLLBACK TRANSACTION;
    -- 可选：抛出错误信息
    RAISERROR('账户A余额不足，转账失败', 16, 1);
END