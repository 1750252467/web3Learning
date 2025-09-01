package logic

import (
	"context"

	"note/week6/zero/test1/api/demo/internal/svc"
	"note/week6/zero/test1/api/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req *types.Request) (resp *types.Response, err error) {
	// 根据传入的name参数构建响应消息
	message := "Hello from " + req.Name
	
	// 返回响应对象
	return &types.Response{
		Message: message,
	}, nil
}
