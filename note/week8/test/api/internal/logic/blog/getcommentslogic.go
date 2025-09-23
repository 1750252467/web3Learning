package blog

import (
	"context"
	"errors"

	"blog/api/api/internal/svc"
	"blog/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取评论列表
func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentsLogic) GetComments(req *types.GetCommentsReq) (resp *types.GetCommentsResp, err error) {
	// 检查请求参数中的帖子ID
	if req.PostId <= 0 {
		return nil, errors.New("帖子ID无效")
	}

	// 查询指定帖子的评论列表，按创建时间倒序排列
	var comments []svc.Comment
	db := l.svcCtx.DB.Model(&svc.Comment{}).
		Where("post_id = ?", req.PostId).
		Order("created_at DESC").
		Find(&comments)

	if db.Error != nil {
		l.Logger.Errorf("查询评论列表失败: %v", db.Error)
		err = db.Error
		return
	}

	// 转换为响应格式
	resp = &types.GetCommentsResp{
		Comments: make([]types.Comment, 0, len(comments)),
	}

	for _, comment := range comments {
		resp.Comments = append(resp.Comments, types.Comment{
			Id:        comment.ID,
			PostId:    comment.PostID,
			UserId:    comment.UserID,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return resp, nil
}
