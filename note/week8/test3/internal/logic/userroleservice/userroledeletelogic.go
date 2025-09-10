package userroleservicelogic

import (
	"context"

	"test3/github.com/example/user"
	"test3/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleDeleteLogic {
	return &UserRoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleDeleteLogic) UserRoleDelete(in *user.UserRoleDeleteReq) (*user.UserRoleDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &user.UserRoleDeleteResp{}, nil
}
