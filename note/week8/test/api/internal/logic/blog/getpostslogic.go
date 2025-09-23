package blog

import (
	"context"

	"blog/api/api/internal/svc"
	"blog/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文章列表
func NewGetPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostsLogic {
	return &GetPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostsLogic) GetPosts(req *types.GetPostsReq) (resp *types.GetPostsResp, err error) {
	// 确定分页参数，默认页码为1，每页大小为10
	page := req.Page
	pageSize := req.Limit
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询帖子列表，按创建时间倒序排列
	var posts []svc.Post
	db := l.svcCtx.DB.Model(&svc.Post{}).
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&posts)

	if db.Error != nil {
		l.Logger.Errorf("查询帖子列表失败: %v", db.Error)
		err = db.Error
		return
	}



	// 转换为响应格式
	resp = &types.GetPostsResp{
		Posts: make([]types.Post, 0, len(posts)),
	}

	for _, post := range posts {
		resp.Posts = append(resp.Posts, types.Post{
			Id:        post.ID,
			UserId:    post.UserID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return resp, nil
}
