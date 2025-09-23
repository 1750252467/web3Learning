package blog

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"

	"blog/api/api/internal/svc"
	"blog/api/api/internal/types"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
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
	// 检查请求参数
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名和密码不能为空")
	}

	// 根据用户名查找用户
	var user svc.User
	err = l.svcCtx.DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		l.Logger.Errorf("用户不存在: %v", err)
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		l.Logger.Errorf("密码验证失败: %v", err)
		return nil, errors.New("用户名或密码错误")
	}

	// 生成JWT令牌
	expiresAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire)).Unix()
	claims := jwt.MapClaims{
		"userId": user.ID,
		"username": user.Username,
		"exp": expiresAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		l.Logger.Errorf("生成令牌失败: %v", err)
		return nil, errors.New("登录失败，请稍后重试")
	}

	// 返回登录成功的信息和令牌
	resp = &types.UserLoginResp{
		Token: tokenString,
	}

	return resp, nil
}
