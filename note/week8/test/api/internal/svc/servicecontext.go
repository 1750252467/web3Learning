package svc

import (
	"blog/api/api/internal/config"
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

// 定义数据模型
// 用户模型
type User struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	Username  string    `gorm:"size:50;not null;uniqueIndex" json:"username"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	Email     string    `gorm:"size:100;uniqueIndex" json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// 文章模型
type Post struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	Title     string    `gorm:"size:200;not null" json:"title"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    int64     `gorm:"index" json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// 评论模型
type Comment struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	UserID    int64     `gorm:"index" json:"userId"`
	PostID    int64     `gorm:"index" json:"postId"`
	CreatedAt time.Time `json:"createdAt"`
}

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	db, err := initDB(c)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 初始化Redis连接
	redisClient := initRedis(c)

	return &ServiceContext{
		Config:      c,
		DB:          db,
		RedisClient: redisClient,
	}
}

// 初始化数据库连接
func initDB(c config.Config) (*gorm.DB, error) {
	// 配置GORM日志
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢SQL阈值
			LogLevel:      logger.Info,   // 日志级别
			Colorful:      true,          // 彩色日志
		},
	)

	// 连接数据库
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// 初始化Redis连接
func initRedis(c config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       0,
	})

	// 测试连接
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Warning: Failed to connect to Redis: %v", err)
	}

	return client
}
