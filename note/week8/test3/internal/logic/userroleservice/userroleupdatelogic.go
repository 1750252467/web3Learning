package userroleservicelogic

import (
	"context"

	"test3/github.com/example/user"
	"test3/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleUpdateLogic {
	return &UserRoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleUpdateLogic) UserRoleUpdate(in *user.UserRoleUpdateReq) (*user.UserRoleUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserRoleUpdateResp{}, nil
}
