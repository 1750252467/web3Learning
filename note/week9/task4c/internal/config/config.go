package config

import (
	"github.com/zeromicro/go-zero/rest"
	"k8s.io/client-go/tools/cache"
)

type Config struct {
	rest.RestConf

	DB struct {
		DataSource string
	}
	Cache cache.Config
}
