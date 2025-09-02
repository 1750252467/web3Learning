package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	DB struct {
		DataSource string
	}

	rest.RestConf
}
