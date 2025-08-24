package controllers

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "note/task4/config"
    "note/task4/models"
)

// CommentController 评论控制器
type CommentController struct {
    DB *gorm.DB
}

// NewCommentController 创建新的评论控制器
func NewCommentController() *CommentController {
    return &CommentController{
        DB: config.DB,
    }
}

// CreateComment 创建评论
func (cc *CommentController) CreateComment(c *gin.Context) {
    // 获取当前用户ID
    userID, exists := c.Request.Context().Value("user_id").(uint)
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    
    // 获取文章ID
    postID, err := strconv.ParseUint(c.Param("post_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }
    
    // 检查文章是否存在
    var post models.Post
    if err := cc.DB.First(&post, postID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }
    
    var comment models.Comment
    
    // 绑定请求体
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // 设置评论相关信息
    comment.UserID = userID
    comment.PostID = uint(postID)
    
    // 创建评论
    if err := cc.DB.Create(&comment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
        return
    }
    
    // 预加载用户信息
    cc.DB.Preload("User").First(&comment, comment.ID)
    
    c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully", "comment": comment})
}

// GetCommentsByPostID 获取某篇文章的所有评论
func (cc *CommentController) GetCommentsByPostID(c *gin.Context) {
    // 获取文章ID
    postID, err := strconv.ParseUint(c.Param("post_id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }
    
    // 检查文章是否存在
    var post models.Post
    if err := cc.DB.First(&post, postID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }
    
    var comments []models.Comment
    
    // 查询该文章的所有评论，并预加载用户信息
    if err := cc.DB.Where("post_id = ?", postID).Preload("User").Find(&comments).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"comments": comments})
}