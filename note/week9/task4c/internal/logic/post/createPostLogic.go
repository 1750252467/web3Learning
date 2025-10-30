package post

import (
	"context"
	"errors"
	"strconv"

	"task4c/internal/model"
	"task4c/internal/svc"
	"task4c/internal/types"

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
	var userID int64
	var parseErr error

	// 首先尝试从context中获取用户ID（正常认证流程）
	ctxUserID, ok := l.ctx.Value("user_id").(int64)
	if ok {
		userID = ctxUserID
		l.Logger.Info("Using user_id from context", logx.Field("user_id", userID))
	} else {
		// 尝试类型转换，处理可能的类型不匹配
		if userIDFloat, ok := l.ctx.Value("user_id").(float64); ok {
			userID = int64(userIDFloat)
			l.Logger.Info("Converted float user_id to int64", logx.Field("user_id", userID))
		} else if userIDStr, ok := l.ctx.Value("user_id").(string); ok {
			// 尝试将字符串类型的user_id转换为int64
			if userID, parseErr = strconv.ParseInt(userIDStr, 10, 64); parseErr == nil {
				l.Logger.Info("Converted string user_id from context to int64", logx.Field("user_id", userID))
			}
		}
	}

	// 如果从context中没有获取到有效的user_id，尝试从请求体中获取（用于开发测试）
	if userID <= 0 && req.UserId != "" {
		if userID, parseErr = strconv.ParseInt(req.UserId, 10, 64); parseErr != nil {
			l.Logger.Error("Failed to parse user_id from request body", logx.Field("error", parseErr.Error()))
			return nil, errors.New("invalid user_id format")
		}
		l.Logger.Info("Using user_id from request body", logx.Field("user_id", userID))
	}

	// 验证用户ID
	if userID <= 0 {
		l.Logger.Error("No valid user_id found in context or request body")
		return nil, errors.New("unauthorized: user not authenticated and no valid user_id provided")
	}

	// 验证请求参数
	if req.Title == "" || req.Content == "" {
		return nil, errors.New("title and content cannot be empty")
	}

	// 创建文章数据
	post := &model.Posts{
		Title:   req.Title,
		Content: req.Content,
		UserId:  userID,
	}

	// 插入文章到数据库
	result, err := l.svcCtx.PostModel.Insert(l.ctx, post)
	if err != nil {
		l.Logger.Error("Failed to insert post", logx.Field("error", err.Error()))
		return nil, errors.New("failed to create post")
	}

	// 获取插入的文章ID
	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Error("Failed to get last insert id", logx.Field("error", err.Error()))
		return nil, errors.New("failed to create post")
	}

	// 返回响应
	return &types.CreatePostResp{
		Id: id,
	}, nil
}
