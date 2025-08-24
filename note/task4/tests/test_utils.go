package tests

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "note/task4/config"
    "note/task4/models"
    "note/task4/utils"
)

// setupTestDB 设置测试数据库
func setupTestDB() *gorm.DB {
    // 使用SQLite内存数据库进行测试
    db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    // 迁移表结构
    err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
    if err != nil {
        panic("Failed to migrate database")
    }

    return db
}

// setupTestRouter 设置测试路由
func setupTestRouter(db *gorm.DB) *gin.Engine {
    // 替换全局DB
    config.DB = db

    // 设置路由
    r := gin.Default()

    return r
}

// createTestUser 创建测试用户
func createTestUser(db *gorm.DB) models.User {
    passwordHash, _ := utils.HashPassword("password123")
    user := models.User{
        Username: "testuser",
        Password: passwordHash,
        Email:    "test@example.com",
    }
    db.Create(&user)
    return user
}

// createTestPost 创建测试文章
func createTestPost(db *gorm.DB, userID uint) models.Post {
    post := models.Post{
        Title:   "Test Post",
        Content: "Test Content",
        UserID:  userID,
    }
    db.Create(&post)
    return post
}

// TestAuthMiddleware 测试环境下的认证中间件
func TestAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 测试环境下的认证中间件
        userID := c.Request.Header.Get("X-Test-User-ID")
        if userID != "" {
            // 使用 gin.Context 的 WithValue 方法
            c.Set("user_id", uint(1))
            c.Next()
            return
        }
        c.JSON(401, gin.H{"error": "Unauthorized"})
        c.Abort()
    }
}