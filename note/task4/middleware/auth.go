package middleware

import (
    "context"
    "net/http"
    "strings"
    
    "github.com/gin-gonic/gin"
    _ "github.com/golang-jwt/jwt/v5"
    "note/task4/utils"
)

// JWTAuthMiddleware JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 从请求头中获取Authorization字段
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }
        
        // 检查Authorization格式
        parts := strings.SplitN(authHeader, " ", 2)
        if !(len(parts) == 2 && parts[0] == "Bearer") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
            c.Abort()
            return
        }
        
        // 验证令牌
        tokenString := parts[1]
        claims, err := utils.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }
        
        // 将用户信息添加到上下文
        userID := uint((*claims)["user_id"].(float64))
        username := (*claims)["username"].(string)
        
        // 创建新的上下文，添加用户信息
        ctx := context.WithValue(c.Request.Context(), "user_id", userID)
        ctx = context.WithValue(ctx, "username", username)
        
        // 替换请求上下文
        c.Request = c.Request.WithContext(ctx)
        
        c.Next()
    }
}