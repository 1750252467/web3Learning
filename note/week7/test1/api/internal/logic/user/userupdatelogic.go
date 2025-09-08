package user

import (
	"context"
	"database/sql"
	"errors"

	"note/week7/test1/api/internal/svc"
	"note/week7/test1/api/internal/types"
	"note/week7/test1/api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户信息
func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {
	// 查找用户是否存在
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询数据失败")
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// 更新用户昵称
	user.Nickname = sql.NullString{String: req.Nickname, Valid: true}
	if err := l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
		return nil, errors.New("更新用户信息失败")
	}

	return &types.UserUpdateResp{
		Flag: true,
	}, nil
}
