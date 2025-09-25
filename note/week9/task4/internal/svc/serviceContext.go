package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"task4/internal/config"
	"task4/model"
)

type ServiceContext struct {
	Config       config.Config
	UserModel    model.UsersModel
	PostModel    model.PostsModel
	CommentModel model.CommentsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		UserModel:    model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		PostModel:    model.NewPostsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		CommentModel: model.NewCommentsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		Config:       c,
	}
}
