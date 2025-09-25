package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	_ UsersModel = (*customUsersModel)(nil)
	// 添加用户名缓存前缀
	cacheZeroDemoUsersUsernamePrefix = "cache:zeroDemo:users:username:"
)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		// 添加通过用户名查找用户的方法

	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}

// FindOneByUsername 通过用户名查找用户
