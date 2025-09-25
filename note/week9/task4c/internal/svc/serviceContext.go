package svc

import (
	"task4c/internal/config"
	"task4c/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 为了解决 undefined: model 的问题，需要导入 model

	UsersModel   model.UsersModel
	PostModel    model.PostsModel
	CommentModel model.CommentsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		// 为解决 undefined: sqlx 的问题，需要导入 sqlx 包，此处先假定 sqlx 包已正确导入，给出修正后的代码
		UsersModel: model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),

		PostModel:    model.NewPostsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		CommentModel: model.NewCommentsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
