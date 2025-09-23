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

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建评论
func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.CreateCommentResp, err error) {
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
	if req.Content == "" {
		return nil, errors.New("评论内容不能为空")
	}

	// 检查帖子是否存在
	var post svc.Post
	db := l.svcCtx.DB.Where("id = ?", req.PostId).First(&post)
	if db.Error != nil {
		if db.Error.Error() == "record not found" {
			return nil, errors.New("帖子不存在")
		}
		l.Logger.Errorf("查询帖子失败: %v", db.Error)
		err = db.Error
		return
	}

	// 创建新评论
	newComment := svc.Comment{
		PostID:    req.PostId,
		UserID:    userId,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}

	db = l.svcCtx.DB.Create(&newComment)
	if db.Error != nil {
		l.Logger.Errorf("创建评论失败: %v", db.Error)
		err = db.Error
		return
	}

	// 返回创建成功的评论信息
	resp = &types.CreateCommentResp{
		Id: newComment.ID,
	}

	return resp, nil
}
