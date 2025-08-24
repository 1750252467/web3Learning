package config

import (
    "database/sql"
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
    // 数据库配置
    username := "root"
    password := "080657"
    host := "localhost"
    port := "3306"
    dbName := "web3Learning"
    
    // 首先尝试连接到MySQL服务器（不指定数据库）
    dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port)
    sqlDB, err := sql.Open("mysql", dsnWithoutDB)
    if err != nil {
        log.Fatalf("Failed to connect to MySQL server: %v", err)
    }
    defer sqlDB.Close()
    
    // 创建数据库
    _, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
    if err != nil {
        log.Fatalf("Failed to create database: %v", err)
    }
    
    // 连接到新创建的数据库
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    
    // 自动迁移表结构
    // Migrate will create tables, missing foreign keys, constraints, columns and indexes.
    // It will not delete/change existing columns and their types
    // AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}