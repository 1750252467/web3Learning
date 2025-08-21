package tests

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "note/task4/config"
    "note/task4/controllers"
    "note/task4/models"
    "note/task4/utils"
)

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

func setupTestRouter(db *gorm.DB) *gin.Engine {
    // 替换全局DB
    config.DB = db

    // 设置路由
    r := gin.Default()

    // 公共路由
    public := r.Group("/")
    {
        // 用户注册和登录
        userController := controllers.NewUserController()
        public.POST("/register", userController.Register)
        public.POST("/login", userController.Login)

        // 文章相关路由
        posts := public.Group("/posts")
        {
            // 文章评论路由
            comments := posts.Group("/:post_id/comments")
            {
                commentController := controllers.NewCommentController()
                comments.GET("", commentController.GetCommentsByPostID)
            }
        }
    }

    // 受保护的路由
    protected := r.Group("/")
    protected.Use(func(c *gin.Context) {
        // 测试环境下的认证中间件
        userID := c.Request.Header.Get("X-Test-User-ID")
        if userID != "" {
            c.Request = c.Request.WithContext(gin.ContextWithValue(c.Request.Context(), "user_id", uint(1)))
            c.Next()
            return
        }
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        c.Abort()
    })
    {
        // 评论管理使用嵌套路由
        posts := protected.Group("/posts")
        {
            comments := posts.Group("/:post_id/comments")
            {
                commentController := controllers.NewCommentController()
                comments.POST("", commentController.CreateComment)
            }
        }
    }

    return r
}

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

func createTestPost(db *gorm.DB, userID uint) models.Post {
    post := models.Post{
        Title:   "Test Post",
        Content: "Test Content",
        UserID:  userID,
    }
    db.Create(&post)
    return post
}

func TestCreateComment(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)
    post := createTestPost(db, user.ID)

    // 测试用例1: 成功创建评论
    commentJSON := `{
        "content": "This is a test comment"
    }`

    req, _ := http.NewRequest("POST", "/posts/"+string(rune(post.ID))+"/comments", strings.NewReader(commentJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)

    // 验证评论是否创建成功
    var comment models.Comment
    result := db.Where("post_id = ?", post.ID).First(&comment)
    assert.NoError(t, result.Error)
    assert.Equal(t, "This is a test comment", comment.Content)
    assert.Equal(t, user.ID, comment.UserID)
    assert.Equal(t, post.ID, comment.PostID)

    // 测试用例2: 未授权
    req, _ = http.NewRequest("POST", "/posts/"+string(rune(post.ID))+"/comments", strings.NewReader(commentJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)

    // 测试用例3: 文章不存在
    req, _ = http.NewRequest("POST", "/posts/999/comments", strings.NewReader(commentJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)

    // 测试用例4: 无效的请求体
    invalidJSON := `{}`  // 缺少内容

    req, _ = http.NewRequest("POST", "/posts/"+string(rune(post.ID))+"/comments", strings.NewReader(invalidJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetCommentsByPostID(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)
    post := createTestPost(db, user.ID)

    // 创建测试评论
    comment1 := models.Comment{
        Content: "Comment 1",
        UserID:  user.ID,
        PostID:  post.ID,
    }
    comment2 := models.Comment{
        Content: "Comment 2",
        UserID:  user.ID,
        PostID:  post.ID,
    }
    db.Create(&comment1)
    db.Create(&comment2)

    // 测试用例1: 成功获取评论列表
    req, _ := http.NewRequest("GET", "/posts/"+string(rune(post.ID))+"/comments", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证响应
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    comments, ok := response["comments"].([]interface{})
    assert.True(t, ok)
    assert.Equal(t, 2, len(comments))

    // 测试用例2: 文章不存在
    req, _ = http.NewRequest("GET", "/posts/999/comments", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)

    // 测试用例3: 无效的ID格式
    req, _ = http.NewRequest("GET", "/posts/invalid/comments", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}