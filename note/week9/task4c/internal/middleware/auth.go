package middleware

import (
	"context"
	"net/http"
	"strings"

	"task4c/internal/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// JWTAuthMiddleware 创建JWT认证中间件
func JWTAuthMiddleware(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 从请求头中获取Authorization
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				httpx.Error(w, http.ErrNoCookie)
				return
			}

			// 检查Authorization格式
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				httpx.Error(w, http.ErrNoCookie)
				return
			}

			// 解析JWT Token
			tokenString := parts[1]
			claims := &jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.Auth.AccessSecret), nil
			})

			if err != nil || !token.Valid {
				logx.Error("Invalid token", logx.Field("error", err))
				httpx.Error(w, http.ErrNoCookie)
				return
			}

			// 从claims中获取用户信息
			userID, ok := (*claims)["userId"].(float64)
			if !ok {
				logx.Error("Invalid userId in token")
				httpx.Error(w, http.ErrNoCookie)
				return
			}

			// 将用户信息添加到context
			ctx := context.WithValue(r.Context(), "user_id", int64(userID))
			if username, ok := (*claims)["username"].(string); ok {
				ctx = context.WithValue(ctx, "username", username)
			}

			// 继续处理请求
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
