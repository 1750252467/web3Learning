package user

import (
	"context"

	"note/week6/zero/test2/api/internal/svc"
	"note/week6/zero/test2/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	m := map[int64]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}
	nickname := "unknow"

	if name, ok := m[req.UserId]; ok {
		nickname = name
	}

	return &types.UserInfoResp{
		UserId:   req.UserId,
		Nickname: nickname,
	}, nil
}
