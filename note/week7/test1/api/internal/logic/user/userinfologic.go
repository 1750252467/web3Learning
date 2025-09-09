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
	// 添加调试日志
	l.Logger.Info("Processing UserInfo request", logx.Field("userId", req.UserId))
	
	// 使用新添加的FindOneByUserId方法，按user_id字段查询用户
	user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, req.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			l.Logger.Info("User not found", logx.Field("userId", req.UserId))
			return nil, errors.New("user not found")
		}
		l.Logger.Error("Database query failed", logx.Field("error", err.Error()), logx.Field("userId", req.UserId))
		return nil, errors.New("查询数据失败")
	}
	l.Logger.Info("User found", logx.Field("userId", req.UserId), logx.Field("dbUserId", user.UserId), logx.Field("nickname", user.Nickname.String))

	return &types.UserInfoResp{
		UserId:   user.UserId, // 返回实际的userId而不是主键id
		Nickname: user.Nickname.String,
	}, nil
}
