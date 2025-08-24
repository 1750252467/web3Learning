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

func TestUserRegister(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)

    // 注册测试路由
    userController := controllers.NewUserController()
    r.POST("/register", userController.Register)
    r.POST("/login", userController.Login)

    // 测试用例1: 成功注册
    userJSON := `{
        "username": "testuser",
        "password": "password123",
        "email": "test@example.com"
    }`

    req, _ := http.NewRequest("POST", "/register", strings.NewReader(userJSON))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)

    // 验证用户是否创建成功
    var user models.User
    result := db.Where("username = ?", "testuser").First(&user)
    assert.NoError(t, result.Error)
    assert.Equal(t, "testuser", user.Username)
    assert.Equal(t, "test@example.com", user.Email)
    assert.NotEmpty(t, user.Password)  // 密码应该被加密

    // 测试用例2: 用户名已存在
    req, _ = http.NewRequest("POST", "/register", strings.NewReader(userJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusConflict, w.Code)

    // 测试用例3: 无效的请求体
    invalidJSON := `{
        "username": "testuser2",
        "email": "test2@example.com"
    }`  // 缺少密码

    req, _ = http.NewRequest("POST", "/register", strings.NewReader(invalidJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserLogin(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)

    // 注册测试路由
    userController := controllers.NewUserController()
    r.POST("/register", userController.Register)
    r.POST("/login", userController.Login)

    // 先创建一个测试用户
    user := createTestUser(db)

    // 测试用例1: 成功登录
    loginJSON := `{
        "username": "testuser",
        "password": "password123"
    }`

    req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginJSON))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证响应中是否包含token
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Contains(t, response, "token")
    assert.Equal(t, float64(user.ID), response["user_id"])
    assert.Equal(t, "testuser", response["username"])

    // 测试用例2: 无效的用户名
    invalidLoginJSON := `{
        "username": "nonexistent",
        "password": "password123"
    }`

    req, _ = http.NewRequest("POST", "/login", strings.NewReader(invalidLoginJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)

    // 测试用例3: 无效的密码
    invalidPasswordJSON := `{
        "username": "testuser",
        "password": "wrongpassword"
    }`

    req, _ = http.NewRequest("POST", "/login", strings.NewReader(invalidPasswordJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)
}


func TestUserRegister(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)

    // 测试用例1: 成功注册
    userJSON := `{
        "username": "testuser",
        "password": "password123",
        "email": "test@example.com"
    }`

    req, _ := http.NewRequest("POST", "/register", strings.NewReader(userJSON))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)

    // 验证用户是否创建成功
    var user models.User
    result := db.Where("username = ?", "testuser").First(&user)
    assert.NoError(t, result.Error)
    assert.Equal(t, "testuser", user.Username)
    assert.Equal(t, "test@example.com", user.Email)
    assert.NotEmpty(t, user.Password)  // 密码应该被加密

    // 测试用例2: 用户名已存在
    req, _ = http.NewRequest("POST", "/register", strings.NewReader(userJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusConflict, w.Code)

    // 测试用例3: 无效的请求体
    invalidJSON := `{
        "username": "testuser2",
        "email": "test2@example.com"
    }`  // 缺少密码

    req, _ = http.NewRequest("POST", "/register", strings.NewReader(invalidJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserLogin(t *testing.T) {
    // 设置测试环境
    db := setupTestDB()
    r := setupTestRouter(db)

    // 先创建一个测试用户
    user := models.User{
        Username: "testuser",
        Password: "$2a$14$8i4Kz6Fv5O3nUY9L4c5D..QeA5R7X7X7X7X7X7X7X7X7X7X7X7", // 密码: password123
        Email:    "test@example.com",
    }
    db.Create(&user)

    // 测试用例1: 成功登录
    loginJSON := `{
        "username": "testuser",
        "password": "password123"
    }`

    req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginJSON))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    // 验证响应中是否包含token
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Contains(t, response, "token")
    assert.Equal(t, float64(user.ID), response["user_id"])
    assert.Equal(t, "testuser", response["username"])

    // 测试用例2: 无效的用户名
    invalidLoginJSON := `{
        "username": "nonexistent",
        "password": "password123"
    }`

    req, _ = http.NewRequest("POST", "/login", strings.NewReader(invalidLoginJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)

    // 测试用例3: 无效的密码
    invalidPasswordJSON := `{
        "username": "testuser",
        "password": "wrongpassword"
    }`

    req, _ = http.NewRequest("POST", "/login", strings.NewReader(invalidPasswordJSON))
    req.Header.Set("Content-Type", "application/json")
    w = httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, http.StatusUnauthorized, w.Code)
}