package routes

import (
    "github.com/gin-gonic/gin"
    "note/task4/controllers"
    "note/task4/middleware"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
    r := gin.Default()
    
    // 公共路由
    public := r.Group("/")
    {
        // 用户注册和登录
        userController := controllers.NewUserController()
        public.POST("/register", userController.Register)
        public.POST("/login", userController.Login)
        
        // 文章相关路由
        posts := public.Group("/posts")
        {
            // 获取文章列表
            postController := controllers.NewPostController()
            posts.GET("", postController.GetAllPosts)
            
            // 文章评论路由
            comments := posts.Group("/:post_id/comments")
            {
                commentController := controllers.NewCommentController()
                comments.GET("", commentController.GetCommentsByPostID)
            }
            
            // 获取单篇文章
            posts.GET("/details/:id", postController.GetPost)
        }
    }
    
    // 受保护的路由
    protected := r.Group("/")
    protected.Use(middleware.JWTAuthMiddleware())
    {
        // 文章管理
        postController := controllers.NewPostController()
        protected.POST("/posts", postController.CreatePost)
        protected.PUT("/posts/:id", postController.UpdatePost)
        protected.DELETE("/posts/:id", postController.DeletePost)
        
        // 评论管理使用嵌套路由
          posts := protected.Group("/posts")
          {
              comments := posts.Group("/:post_id/comments")
              {
                  commentController := controllers.NewCommentController()
                  comments.POST("", commentController.CreateComment)
              }
          }
    }
    
    return r
}