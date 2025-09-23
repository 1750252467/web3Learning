package blog

import (
	"context"
	"errors"

	"blog/api/api/internal/svc"
	"blog/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文章详情
func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostLogic) GetPost(req *types.GetPostReq) (resp *types.GetPostResp, err error) {
	// 根据ID查询帖子
	var post svc.Post
	db := l.svcCtx.DB.Where("id = ?", req.Id).First(&post)

	if db.Error != nil {
		// 检查是否是记录未找到的错误
		if db.Error.Error() == "record not found" {
			return nil, errors.New("帖子不存在")
		}
		l.Logger.Errorf("查询帖子详情失败: %v", db.Error)
		err = db.Error
		return
	}

	// 转换为响应格式
	resp = &types.GetPostResp{
		Id:        post.ID,
		UserId:    post.UserID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return resp, nil
}
