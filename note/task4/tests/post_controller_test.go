package tests

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "note/task4/controllers"
    "note/task4/models"
)

func TestCreatePost(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)

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
            // 获取文章列表
            postController := controllers.NewPostController()
            posts.GET("", postController.GetAllPosts)

            // 获取单篇文章
            posts.GET("/details/:id", postController.GetPost)
        }
    }

    // 受保护的路由
    protected := r.Group("/")
    protected.Use(TestAuthMiddleware())
    {
        // 文章管理
        postController := controllers.NewPostController()
        protected.POST("/posts", postController.CreatePost)
        protected.PUT("/posts/:id", postController.UpdatePost)
        protected.DELETE("/posts/:id", postController.DeletePost)
    }

    // 测试用例1: 成功创建文章
    postJSON := `{
        "title": "Test Post",
        "content": "This is a test post content"
    }`

    req, _ := http.NewRequest("POST", "/posts", strings.NewReader(postJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)

    // 验证文章是否创建成功
    var post models.Post
    result := db.Where("title = ?", "Test Post").First(&post)
    assert.NoError(t, result.Error)
    assert.Equal(t, "Test Post", post.Title)
    assert.Equal(t, "This is a test post content", post.Content)
    assert.Equal(t, user.ID, post.UserID)

    // 测试用例2: 未授权
    req, _ = http.NewRequest("POST", "/posts", strings.NewReader(postJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)

    // 测试用例3: 无效的请求体
    invalidJSON := `{
        "title": "Invalid Post"
    }`  // 缺少内容

    req, _ = http.NewRequest("POST", "/posts", strings.NewReader(invalidJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetAllPosts(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)

    // 公共路由
    public := r.Group("/")
    {
        // 文章相关路由
        posts := public.Group("/posts")
        {
            // 获取文章列表
            postController := controllers.NewPostController()
            posts.GET("", postController.GetAllPosts)
        }
    }

    // 创建测试文章
    post1 := models.Post{
        Title:   "Post 1",
        Content: "Content 1",
        UserID:  user.ID,
    }
    post2 := models.Post{
        Title:   "Post 2",
        Content: "Content 2",
        UserID:  user.ID,
    }
    db.Create(&post1)
    db.Create(&post2)

    // 测试获取所有文章
    req, _ := http.NewRequest("GET", "/posts", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证响应
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    posts, ok := response["posts"].([]interface{})
    assert.True(t, ok)
    assert.Equal(t, 2, len(posts))
}

func TestGetPost(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)

    // 公共路由
    public := r.Group("/")
    {
        // 文章相关路由
        posts := public.Group("/posts")
        {
            // 获取单篇文章
            postController := controllers.NewPostController()
            posts.GET("/details/:id", postController.GetPost)
        }
    }

    // 创建测试文章
    post := models.Post{
        Title:   "Test Post",
        Content: "Test Content",
        UserID:  user.ID,
    }
    db.Create(&post)

    // 测试用例1: 成功获取文章
    req, _ := http.NewRequest("GET", "/posts/details/"+string(rune(post.ID)), nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证响应
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    postData, ok := response["post"].(map[string]interface{})
    assert.True(t, ok)
    assert.Equal(t, float64(post.ID), postData["ID"])
    assert.Equal(t, "Test Post", postData["title"])
    assert.Equal(t, "Test Content", postData["content"])

    // 测试用例2: 文章不存在
    req, _ = http.NewRequest("GET", "/posts/details/999", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)

    // 测试用例3: 无效的ID格式
    req, _ = http.NewRequest("GET", "/posts/details/invalid", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdatePost(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)

    // 受保护的路由
    protected := r.Group("/")
    protected.Use(TestAuthMiddleware())
    {
        // 文章管理
        postController := controllers.NewPostController()
        protected.PUT("/posts/:id", postController.UpdatePost)
    }

    // 创建测试文章
    post := models.Post{
        Title:   "Original Title",
        Content: "Original Content",
        UserID:  user.ID,
    }
    db.Create(&post)

    // 测试用例1: 成功更新文章
    updateJSON := `{
        "title": "Updated Title",
        "content": "Updated Content"
    }`

    req, _ := http.NewRequest("PUT", "/posts/"+string(rune(post.ID)), strings.NewReader(updateJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证文章是否更新成功
    var updatedPost models.Post
    db.First(&updatedPost, post.ID)
    assert.Equal(t, "Updated Title", updatedPost.Title)
    assert.Equal(t, "Updated Content", updatedPost.Content)

    // 测试用例2: 无权限更新
    anotherUser := models.User{
        Username: "anotheruser",
        Password: "$2a$14$8i4Kz6Fv5O3nUY9L4c5D..QeA5R7X7X7X7X7X7X7X7X7X7X7X7", // password123
        Email:    "another@example.com",
    }
    db.Create(&anotherUser)

    req, _ = http.NewRequest("PUT", "/posts/"+string(rune(post.ID)), strings.NewReader(updateJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(anotherUser.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusForbidden, w.Code)

    // 测试用例3: 文章不存在
    req, _ = http.NewRequest("PUT", "/posts/999", strings.NewReader(updateJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeletePost(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)

    // 受保护的路由
    protected := r.Group("/")
    protected.Use(TestAuthMiddleware())
    {
        // 文章管理
        postController := controllers.NewPostController()
        protected.DELETE("/posts/:id", postController.DeletePost)
    }

    // 创建测试文章
    post := models.Post{
        Title:   "Test Post",
        Content: "Test Content",
        UserID:  user.ID,
    }
    db.Create(&post)

    // 测试用例1: 成功删除文章
    req, _ := http.NewRequest("DELETE", "/posts/"+string(rune(post.ID)), nil)
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证文章是否删除成功
    var deletedPost models.Post
    result := db.First(&deletedPost, post.ID)
    assert.Error(t, result.Error)

    // 测试用例2: 无权限删除
    anotherUser := models.User{
        Username: "anotheruser",
        Password: "$2a$14$8i4Kz6Fv5O3nUY9L4c5D..QeA5R7X7X7X7X7X7X7X7X7X7X7X7", // password123
        Email:    "another@example.com",
    }
    db.Create(&anotherUser)

    // 重新创建文章
    db.Create(&post)

    req, _ = http.NewRequest("DELETE", "/posts/"+string(rune(post.ID)), nil)
    req.Header.Set("X-Test-User-ID", string(rune(anotherUser.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusForbidden, w.Code)

    // 测试用例3: 文章不存在
    req, _ = http.NewRequest("DELETE", "/posts/999", nil)
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)
}


func TestCreatePost(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)

    // 测试用例1: 成功创建文章
    postJSON := `{
        "title": "Test Post",
        "content": "This is a test post content"
    }`

    req, _ := http.NewRequest("POST", "/posts", strings.NewReader(postJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)

    // 验证文章是否创建成功
    var post models.Post
    result := db.Where("title = ?", "Test Post").First(&post)
    assert.NoError(t, result.Error)
    assert.Equal(t, "Test Post", post.Title)
    assert.Equal(t, "This is a test post content", post.Content)
    assert.Equal(t, user.ID, post.UserID)

    // 测试用例2: 未授权
    req, _ = http.NewRequest("POST", "/posts", strings.NewReader(postJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)

    // 测试用例3: 无效的请求体
    invalidJSON := `{
        "title": "Invalid Post"
    }`  // 缺少内容

    req, _ = http.NewRequest("POST", "/posts", strings.NewReader(invalidJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetAllPosts(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)

    // 创建测试文章
    post1 := models.Post{
        Title:   "Post 1",
        Content: "Content 1",
        UserID:  user.ID,
    }
    post2 := models.Post{
        Title:   "Post 2",
        Content: "Content 2",
        UserID:  user.ID,
    }
    db.Create(&post1)
    db.Create(&post2)

    // 测试获取所有文章
    req, _ := http.NewRequest("GET", "/posts", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证响应
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    posts, ok := response["posts"].([]interface{})
    assert.True(t, ok)
    assert.Equal(t, 2, len(posts))
}

func TestGetPost(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)

    // 创建测试文章
    post := models.Post{
        Title:   "Test Post",
        Content: "Test Content",
        UserID:  user.ID,
    }
    db.Create(&post)

    // 测试用例1: 成功获取文章
    req, _ := http.NewRequest("GET", "/posts/details/"+string(rune(post.ID)), nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证响应
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    postData, ok := response["post"].(map[string]interface{})
    assert.True(t, ok)
    assert.Equal(t, float64(post.ID), postData["ID"])
    assert.Equal(t, "Test Post", postData["title"])
    assert.Equal(t, "Test Content", postData["content"])

    // 测试用例2: 文章不存在
    req, _ = http.NewRequest("GET", "/posts/details/999", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)

    // 测试用例3: 无效的ID格式
    req, _ = http.NewRequest("GET", "/posts/details/invalid", nil)
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdatePost(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)

    // 创建测试文章
    post := models.Post{
        Title:   "Original Title",
        Content: "Original Content",
        UserID:  user.ID,
    }
    db.Create(&post)

    // 测试用例1: 成功更新文章
    updateJSON := `{
        "title": "Updated Title",
        "content": "Updated Content"
    }`

    req, _ := http.NewRequest("PUT", "/posts/"+string(rune(post.ID)), strings.NewReader(updateJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证文章是否更新成功
    var updatedPost models.Post
    db.First(&updatedPost, post.ID)
    assert.Equal(t, "Updated Title", updatedPost.Title)
    assert.Equal(t, "Updated Content", updatedPost.Content)

    // 测试用例2: 无权限更新
    anotherUser := models.User{
        Username: "anotheruser",
        Password: "$2a$14$8i4Kz6Fv5O3nUY9L4c5D..QeA5R7X7X7X7X7X7X7X7X7X7X7X7", // password123
        Email:    "another@example.com",
    }
    db.Create(&anotherUser)

    req, _ = http.NewRequest("PUT", "/posts/"+string(rune(post.ID)), strings.NewReader(updateJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(anotherUser.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusForbidden, w.Code)

    // 测试用例3: 文章不存在
    req, _ = http.NewRequest("PUT", "/posts/999", strings.NewReader(updateJSON))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDeletePost(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)
    user := createTestUser(db)

    // 创建测试文章
    post := models.Post{
        Title:   "Test Post",
        Content: "Test Content",
        UserID:  user.ID,
    }
    db.Create(&post)

    // 测试用例1: 成功删除文章
    req, _ := http.NewRequest("DELETE", "/posts/"+string(rune(post.ID)), nil)
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证文章是否删除成功
    var deletedPost models.Post
    result := db.First(&deletedPost, post.ID)
    assert.Error(t, result.Error)

    // 测试用例2: 无权限删除
    anotherUser := models.User{
        Username: "anotheruser",
        Password: "$2a$14$8i4Kz6Fv5O3nUY9L4c5D..QeA5R7X7X7X7X7X7X7X7X7X7X7X7", // password123
        Email:    "another@example.com",
    }
    db.Create(&anotherUser)

    // 重新创建文章
    db.Create(&post)

    req, _ = http.NewRequest("DELETE", "/posts/"+string(rune(post.ID)), nil)
    req.Header.Set("X-Test-User-ID", string(rune(anotherUser.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusForbidden, w.Code)

    // 测试用例3: 文章不存在
    req, _ = http.NewRequest("DELETE", "/posts/999", nil)
    req.Header.Set("X-Test-User-ID", string(rune(user.ID)))
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)
}