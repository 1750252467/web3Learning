package models

import (
    "gorm.io/gorm"
)

// Comment 模型定义文章评论信息
type Comment struct {
    gorm.Model
    Content string `gorm:"type:text;not null" json:"content"`
    UserID  uint   `gorm:"not null" json:"user_id"`
    User    User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
    PostID  uint   `gorm:"not null" json:"post_id"`
    Post    Post   `gorm:"foreignKey:PostID" json:"post,omitempty"`
}