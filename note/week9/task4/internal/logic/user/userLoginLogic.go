package user

import (
	"context"
	"errors"

	"task4/internal/svc"
	"task4/internal/types"

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
	// todo: add your logic here and delete this line
	// 校验用户是否存在
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	// 校验密码是否正确
	if user.Password != req.Password {
		return nil, errors.New("password not match")
	}
	// 登录成功，返回token
	token, err := l.svcCtx.UserModel.GenerateToken(l.ctx, user.Id, user.Username)
	if err != nil {
		return nil, err
	}
	return &types.UserLoginResp{
		Token: token,
	}, nil
}
