package user

import (
	"context"
	"errors"

	"note/week7/test1/api/internal/svc"
	"note/week7/test1/api/internal/types"
	"note/week7/test1/api/model"

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
	// Convert UserId from int64 to string since that's what the model expects
	// 移除未使用的变量声明
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询数据失败")
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return &types.UserInfoResp{
		UserId:   user.Id,
		Nickname: user.Nickname.String,
	}, nil
}
