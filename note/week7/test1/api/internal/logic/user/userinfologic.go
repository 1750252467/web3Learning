package user

import (
	"context"

	"note/week6/zero/test2/api/internal/svc"
	"note/week6/zero/test2/api/internal/types"
	"note/week6/zero/test2/api/model"

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
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, err.New("查询数据失败")
	}
	if user == nil {
		return nil, err.New("user not found")
	}

	return &types.UserInfoResp{
		UserId:   user.UserId,
		Nickname: user.nickname,
	}, nil
}
