package controllers

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "note/task4/config"
    "note/task4/models"
    "note/task4/utils"
)

// UserController 用户控制器
type UserController struct {
    DB *gorm.DB
}

// NewUserController 创建新的用户控制器
func NewUserController() *UserController {
    return &UserController{
        DB: config.DB,
    }
}

// Register 用户注册
func (uc *UserController) Register(c *gin.Context) {
    var user models.User
    
    // 绑定请求体
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // 检查用户名是否已存在
    var existingUser models.User
    if err := uc.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
        return
    }
    
    // 加密密码
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    
    // 创建新用户
    newUser := models.User{
        Username: user.Username,
        Password: hashedPassword,
        Email:    user.Email,
    }
    
    if err := uc.DB.Create(&newUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user_id": newUser.ID})
}

// Login 用户登录
func (uc *UserController) Login(c *gin.Context) {
    var loginData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    
    // 绑定请求体
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // 查找用户
    var user models.User
    if err := uc.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }
    
    // 验证密码
    if !utils.CheckPasswordHash(loginData.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }
    
    // 生成JWT令牌
    token, err := utils.GenerateToken(user.ID, user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"token": token, "user_id": user.ID, "username": user.Username})
}