package task3

//package main 题目1

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 导入对应数据库驱动
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Employee 员工结构体，与employees表字段对应
type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

// 查询部门为"技术部"的所有员工
func getTechDepartmentEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee

	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	return employees, nil
}

// 查询工资最高的员工
func getHighestSalaryEmployee(db *sqlx.DB) (Employee, error) {
	var employee Employee

	err := db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		return Employee{}, fmt.Errorf("查询失败: %v", err)
	}
	return employee, nil
}

//题目2

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func getExpensiveBooks(db *sqlx.DB) ([]Book, error) {
	var books []Book
	// 执行查询并映射到Book切片，通过类型匹配保证类型安全
	err := db.Select(&books, "SELECT id, title, author, price FROM books WHERE price > ?", 50)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %v", err)
	}
	return books, nil
}

//进阶 gorm
// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

// User 用户模型
type User struct {
	gorm.Model
	Username     string // 用户名
	ArticleCount int    // 文章数量统计
	Posts        []Post // 一对多关联：一个用户有多篇文章
}

// Post 文章模型
type Post struct {
	gorm.Model
	Title         string    // 文章标题
	Content       string    // 文章内容
	UserID        uint      // 外键：关联用户ID
	User          User      // 关联的用户
	CommentCount  int       // 评论数量
	CommentStatus string    // 评论状态："有评论"/"无评论"
	Comments      []Comment // 一对多关联：一篇文章有多条评论
}

// Comment 评论模型
type Comment struct {
	gorm.Model
	Content string // 评论内容
	PostID  uint   // 外键：关联文章ID
	Post    Post   // 关联的文章
}

// 创建模型对应的数据库表
func createTables() error {
	// 数据库连接配置（替换为实际配置）
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("数据库连接失败: %v", err)
	}

	// 自动迁移创建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		return fmt.Errorf("创建表失败: %v", err)
	}
	return nil
}

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

// 查询用户的所有文章及评论
func getUserPostsWithComments(db *gorm.DB, userID uint) (User, error) {
	var user User
	// 使用Preload嵌套预加载关联数据
	result := db.Preload("Posts.Comments").First(&user, userID)
	if result.Error != nil {
		return User{}, fmt.Errorf("查询失败: %v", result.Error)
	}
	return user, nil
}

// 查询评论数量最多的文章
func getMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post
	// 关联查询并按评论数排序
	result := db.Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Select("posts.*, COUNT(comments.id) as comment_count").
		Order("comment_count DESC").
		First(&post)

	if result.Error != nil {
		return Post{}, fmt.Errorf("查询失败: %v", result.Error)
	}
	return post, nil
}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

// 为Post模型添加创建后钩子（AfterCreate）
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 更新用户的文章数量（自增1）
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("article_count", gorm.Expr("article_count + 1")).Error
}

// 为Comment模型添加删除后钩子（AfterDelete）
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var post Post
	// 查询当前评论对应的文章
	if err := tx.First(&post, c.PostID).Error; err != nil {
		return err
	}

	// 查询该文章剩余的评论数量
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount).Error; err != nil {
		return err
	}

	// 如果评论数量为0，更新文章评论状态
	if commentCount == 0 {
		return tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
	}
	return nil
}
