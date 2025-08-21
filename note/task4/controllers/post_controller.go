package controllers

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "note/task4/config"
    "note/task4/models"
)

// PostController 文章控制器
type PostController struct {
    DB *gorm.DB
}

// NewPostController 创建新的文章控制器
func NewPostController() *PostController {
    return &PostController{
        DB: config.DB,
    }
}

// CreatePost 创建文章
func (pc *PostController) CreatePost(c *gin.Context) {
    // 获取当前用户ID
    userID, exists := c.Request.Context().Value("user_id").(uint)
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    
    var post models.Post
    
    // 绑定请求体
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // 设置文章作者
    post.UserID = userID
    
    // 创建文章
    if err := pc.DB.Create(&post).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
        return
    }
    
    c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully", "post": post})
}

// GetAllPosts 获取所有文章
func (pc *PostController) GetAllPosts(c *gin.Context) {
    var posts []models.Post
    
    // 查询所有文章，并预加载用户信息
    if err := pc.DB.Preload("User").Find(&posts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// GetPost 获取单篇文章
func (pc *PostController) GetPost(c *gin.Context) {
    // 获取文章ID
    postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }
    
    var post models.Post
    
    // 查询文章，并预加载用户和评论信息
    if err := pc.DB.Preload("User").Preload("Comments").Preload("Comments.User").First(&post, postID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"post": post})
}

// UpdatePost 更新文章
func (pc *PostController) UpdatePost(c *gin.Context) {
    // 获取当前用户ID
    userID, exists := c.Request.Context().Value("user_id").(uint)
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    
    // 获取文章ID
    postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }
    
    var post models.Post
    
    // 查询文章
    if err := pc.DB.First(&post, postID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }
    
    // 检查是否为文章作者
    if post.UserID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this post"})
        return
    }
    
    // 绑定请求体
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // 更新文章
    if err := pc.DB.Save(&post).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully", "post": post})
}

// DeletePost 删除文章
func (pc *PostController) DeletePost(c *gin.Context) {
    // 获取当前用户ID
    userID, exists := c.Request.Context().Value("user_id").(uint)
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    
    // 获取文章ID
    postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }
    
    var post models.Post
    
    // 查询文章
    if err := pc.DB.First(&post, postID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }
    
    // 检查是否为文章作者
    if post.UserID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this post"})
        return
    }
    
    // 删除文章
    if err := pc.DB.Delete(&post).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}