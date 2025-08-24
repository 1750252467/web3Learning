package models

import (
    "gorm.io/gorm"
)

// User 模型定义用户信息
type User struct {
    gorm.Model
    Username string `gorm:"size:50;not null;unique" json:"username"`
    Password string `gorm:"size:100;not null" json:"-"` // 密码不应该在JSON中显示
    Email    string `gorm:"size:100;unique" json:"email"`
    Posts    []Post `gorm:"foreignKey:UserID" json:"posts,omitempty"`
    Comments []Comment `gorm:"foreignKey:UserID" json:"comments,omitempty"`
}

// BeforeSave 钩子函数，在保存用户前加密密码
func (u *User) BeforeSave(tx *gorm.DB) error {
    // 这里我们会在后续实现密码加密
    return nil
}