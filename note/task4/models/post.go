package models

import (
    "gorm.io/gorm"
)

// Post 模型定义博客文章信息
type Post struct {
    gorm.Model
    Title   string `gorm:"size:100;not null" json:"title"`
    Content string `gorm:"type:text;not null" json:"content"`
    UserID  uint   `gorm:"not null" json:"user_id"`
    User    User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
    Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}