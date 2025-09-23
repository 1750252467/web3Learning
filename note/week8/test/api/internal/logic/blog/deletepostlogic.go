package blog

import (
	"context"
	"errors"
	"strconv"

	"blog/api/api/internal/svc"
	"blog/api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除文章
func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePostLogic) DeletePost(req *types.DeletePostReq) (resp *types.DeletePostResp, err error) {
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

	// 查询帖子
	var post svc.Post
	db := l.svcCtx.DB.Where("id = ?", req.Id).First(&post)

	if db.Error != nil {
		// 检查是否是记录未找到的错误
		if db.Error.Error() == "record not found" {
			return nil, errors.New("帖子不存在")
		}
		l.Logger.Errorf("查询帖子失败: %v", db.Error)
		err = db.Error
		return
	}

	// 检查用户是否有权限删除帖子
	if post.UserID != userId {
		return nil, errors.New("您没有权限删除此帖子")
	}

	// 执行删除操作
	db = l.svcCtx.DB.Delete(&post)
	if db.Error != nil {
		l.Logger.Errorf("删除帖子失败: %v", db.Error)
		err = db.Error
		return
	}

	// 返回删除成功的信息
	resp = &types.DeletePostResp{
		Success: true,
	}

	return resp, nil
}
