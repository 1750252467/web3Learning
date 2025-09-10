package userroleservicelogic

import (
	"context"

	"test3/github.com/example/user"
	"test3/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleListLogic {
	return &UserRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleListLogic) UserRoleList(in *user.UserRoleListReq) (*user.UserRoleListResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserRoleListResp{}, nil
}
