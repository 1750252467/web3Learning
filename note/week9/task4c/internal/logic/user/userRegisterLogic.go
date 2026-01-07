package user

import (
	"context"
	"errors"

	"task4c/internal/model"
	"task4c/internal/svc"
	"task4c/internal/types"

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
	// 1. 参数验证
	if len(req.Username) == 0 || len(req.Password) == 0 || len(req.Email) == 0 {
		return nil, errors.New("用户名、密码和邮箱不能为空")
	}

	// 2. 检查用户名是否已存在
	existingUser, err := l.svcCtx.UsersModel.FindOneByUsername(l.ctx, req.Username)
	if err == nil && existingUser != nil {
		return nil, errors.New("用户名已存在")
	}

	// 3. 检查邮箱是否已存在
	existingEmailUser, err := l.svcCtx.UsersModel.FindOneByEmail(l.ctx, req.Email)
	if err == nil && existingEmailUser != nil {
		return nil, errors.New("邮箱已被注册")
	}

	// 4. 创建新用户
	user := &model.Users{
		Username: req.Username,
		Password: req.Password, // 注意：实际项目中应该对密码进行加密处理
		Email:    req.Email,
	}

	// 插入用户数据
	sqlResult, err := l.svcCtx.UsersModel.Insert(l.ctx, user)
	if err != nil {
		l.Logger.Error("Failed to insert user", logx.Field("username", req.Username), logx.Field("err", err))
		return nil, errors.New("注册失败，请稍后重试")
	}

	// 获取插入的用户ID
	userId, err := sqlResult.LastInsertId()
	if err != nil {
		l.Logger.Error("Failed to get last insert id", logx.Field("err", err))
		return nil, errors.New("注册失败，请稍后重试")
	}

	// 5. 返回注册成功的响应
	resp = &types.UserRegisterResp{
		Id:       userId,
		Username: req.Username,
	}

	l.Logger.Info("User registered successfully", logx.Field("userId", userId), logx.Field("username", req.Username))

	return
}
