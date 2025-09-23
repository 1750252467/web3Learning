package blog

import (
	"context"
	"errors"
	"strconv"
	"time"

	"blog/api/api/internal/svc"
	"blog/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建文章
func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePostLogic) CreatePost(req *types.CreatePostReq) (resp *types.CreatePostResp, err error) {
	// 从上下文中获取用户ID
	userId, ok := l.ctx.Value("userId").(int64)
	if !ok {
		// 尝试从上下文中获取字符串类型的userId
		userIdStr, ok := l.ctx.Value("userId").(string)
		if !ok {
			return nil, errors.New("未授权访问")
		}
		userIdInt, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			return nil, errors.New("用户ID格式错误")
		}
		userId = userIdInt
	}

	// 检查请求参数
	if req.Title == "" || req.Content == "" {
		return nil, errors.New("标题和内容不能为空")
	}

	// 创建新帖子
	newPost := svc.Post{
		UserID:    userId,
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = l.svcCtx.DB.Create(&newPost).Error
	if err != nil {
		l.Logger.Errorf("创建帖子失败: %v", err)
		return nil, errors.New("创建帖子失败，请稍后重试")
	}

	// 返回创建成功的帖子信息
	resp = &types.CreatePostResp{
		Id: newPost.ID,
	}

	return resp, nil
}
