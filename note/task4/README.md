
# Haikus for Codespaces

## 项目结构

创建了以下目录结构来组织代码：

- models/: 数据模型定义
- controllers/: 处理请求的控制器
- routes/: 路由定义
- middleware/: 中间件
- config/: 配置文件
- utils/: 工具函数

## 已实现功能

1. 1.

   **用户认证与授权**

   - 用户注册和登录
   - 密码加密存储
   - JWT 认证中间件

2. 2.

   **文章管理**

   - 创建文章
   - 获取文章列表和详情
   - 更新和删除文章（仅作者可操作）

3. 3.

   **评论功能**

   - 创建评论
   - 获取文章的所有评论

4. 4.

   **数据库配置**

   - 使用 GORM 连接 MySQL 数据库
   - 自动迁移表结构

## 使用说明

1. 1.

   首先修改 `database.go` 中的数据库连接字符串，确保密码和数据库名称正确

2. 2.

   运行以下命令启动服务器：

   ```
   Bash
   运行
   go run main.go
   ```

3. 3.

   服务器将在 [http://localhost:8080](http://localhost:8080/) 启动

## 测试 API

以下是一些常用 API 端点：

- POST /register: 用户注册
- POST /login: 用户登录
- GET /posts: 获取所有文章
- GET /posts/:id: 获取单篇文章
- POST /posts (需要认证): 创建文章
- PUT /posts/:id (需要认证): 更新文章
- DELETE /posts/:id (需要认证): 删除文章
- GET /posts/:post_id/comments: 获取文章评论
- POST /posts/:post_id/comments (需要认证): 创建评论
