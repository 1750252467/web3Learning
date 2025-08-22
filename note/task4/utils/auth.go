package utils

import (
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "time"
)

var jwtSecret = []byte("your-secret-key") // 实际应用中应从环境变量获取

// HashPassword 对密码进行bcrypt加密
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CheckPasswordHash 验证密码是否匹配
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID uint, username string) (string, error) {
    // 设置令牌过期时间
    expirationTime := time.Now().Add(24 * time.Hour)
    
    // 创建声明
    claims := &jwt.MapClaims{
        "user_id":  userID,
        "username": username,
        "exp":      expirationTime.Unix(),
    }
    
    // 创建令牌
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    // 签名令牌
    signedToken, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }
    
    return signedToken, nil
}

// ValidateToken 验证JWT令牌
func ValidateToken(tokenString string) (*jwt.MapClaims, error) {
    // 解析令牌
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    
    if err != nil {
        return nil, err
    }
    
    // 验证令牌
    if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    
    return nil, err
}