package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Host string
		Type string
		Pass string
		Tls  bool
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Log struct {
		Mode                string
		Level               string
		Compress            bool
		KeepDays            int
		StackCoolDownMillis int
	}
	Limiter struct {
		QPS   int
		Burst int
	}
}
