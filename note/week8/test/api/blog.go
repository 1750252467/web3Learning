package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strings"

	"blog/api/api/internal/config"
	"blog/api/api/internal/handler"
	"blog/api/api/internal/svc"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/blog.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 创建自定义中间件配置
	restConf := c.RestConf

	// 创建服务器
	server := rest.MustNewServer(restConf)
	// 注册中间件
	server.Use(jwtAuthMiddleware(c.Auth.AccessSecret))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// JWT认证中间件
func jwtAuthMiddleware(secret string) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 获取令牌
			tokenStr := extractToken(r)
			if tokenStr == "" {
				// 对于不需要认证的路由，直接通过
				path := r.URL.Path
				if path == "/api/v1/user/register" || path == "/api/v1/user/login" || path == "/api/v1/posts" || 
					(path == "/api/v1/post" && r.Method == "GET") || path == "/api/v1/comments" {
					next(w, r)
					return
				}
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("未提供令牌"))
				return
			}

			// 验证令牌
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				// 验证签名算法
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("无效的令牌"))
				return
			}

			// 从令牌中提取用户信息
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("无效的令牌声明"))
				return
			}

			// 将用户信息注入到请求上下文中
			userId, ok := claims["userId"].(float64)
			if ok {
				r = r.WithContext(context.WithValue(r.Context(), "userId", int64(userId)))
			} else if userIdStr, ok := claims["userId"].(string); ok {
				r = r.WithContext(context.WithValue(r.Context(), "userId", userIdStr))
			}

			// 继续处理请求
			next(w, r)
		}
	}
}

// 从请求头中提取令牌
func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		// 格式: Bearer {token}
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
	}
	return ""
}
