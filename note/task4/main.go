package main

import (
    "log"
    "note/task4/config"
    "note/task4/models"
    "note/task4/routes"
)

func main() {
    // 初始化数据库
    config.InitDB()
    
    // 自动迁移表结构
    config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
    
    // 设置路由
    r := routes.SetupRouter()
    
    // 启动服务器
    log.Println("Server starting on port 8080...")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}