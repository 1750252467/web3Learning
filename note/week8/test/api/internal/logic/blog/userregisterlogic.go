package blog

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"

	"blog/api/api/internal/svc"
	"blog/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	// 检查请求参数
	if req.Username == "" || req.Password == "" || req.Email == "" {
		return nil, errors.New("用户名、密码和邮箱不能为空")
	}

	// 检查用户名是否已存在
	var existingUser svc.User
	err = l.svcCtx.DB.Where("username = ?", req.Username).First(&existingUser).Error
	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	err = l.svcCtx.DB.Where("email = ?", req.Email).First(&existingUser).Error
	if err == nil {
		return nil, errors.New("邮箱已被注册")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Logger.Errorf("密码加密失败: %v", err)
		return nil, errors.New("注册失败，请稍后重试")
	}

	// 创建新用户
	newUser := svc.User{
		Username:  req.Username,
		Password:  string(hashedPassword),
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = l.svcCtx.DB.Create(&newUser).Error
	if err != nil {
		l.Logger.Errorf("创建用户失败: %v", err)
		return nil, errors.New("注册失败，请稍后重试")
	}

	// 返回注册成功的用户信息
	resp = &types.UserRegisterResp{
		Id:       newUser.ID,
		Username: newUser.Username,
	}

	return resp, nil
}
