package svc

import (
	"note/week7/test1/api/internal/config"

	"note/week7/test1/api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
