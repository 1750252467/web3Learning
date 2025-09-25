package user

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"task4c/internal/model"
	"task4c/internal/svc"
	"task4c/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// 自定义Claims结构体
type CustomClaims struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const (
	// 假设使用的Token密钥
	TokenSecret = "your-jwt-secret-key"
	// Token过期时间
	TokenExpire = 7 * 24 * time.Hour
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// 1. 检查请求参数
	if len(req.Username) == 0 || len(req.Password) == 0 {
		err = errors.New("用户名或密码不能为空")
		l.Logger.Error("Invalid login request", logx.Field("username", req.Username))
		return
	}

	// 2. 通过用户名查找用户
	user, err := l.svcCtx.UsersModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		// 处理用户不存在的情况
		if err == model.ErrNotFound {
			err = errors.New("用户不存在")
		} else {
			l.Logger.Error("Failed to find user by username", logx.Field("username", req.Username), logx.Field("err", err))
			err = errors.New("登录失败，请稍后重试")
		}
		return
	}

	// 3. 验证密码（注意：实际项目中应该使用加密存储的密码进行验证）
	if user.Password != req.Password {
		err = errors.New("密码错误")
		l.Logger.Error("Password mismatch", logx.Field("username", req.Username))
		return
	}

	// 4. 生成JWT Token
	expirationTime := time.Now().Add(TokenExpire)
	claims := &CustomClaims{
		UserId:   user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// 创建Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(TokenSecret))
	if err != nil {
		l.Logger.Error("Failed to generate token", logx.Field("userId", user.Id), logx.Field("err", err))
		err = errors.New("登录失败，请稍后重试")
		return
	}

	// 5. 返回登录成功的响应
	resp = &types.UserLoginResp{
		Token: tokenString,
	}

	l.Logger.Info("User login successful", logx.Field("userId", user.Id), logx.Field("username", user.Username))

	return
}
