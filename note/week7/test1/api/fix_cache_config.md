# 修复 c.Cache 未定义错误

## 问题分析
编译错误 `internal/svc/servicecontext.go:18:67: c.Cache undefined (type config.Config has no field or method Cache)` 表明在配置结构体中缺少 Cache 字段，但在初始化 UserModel 时尝试使用该字段。

## 修复步骤

### 1. 修改 config.go 文件，添加 Cache 配置

```go
// 修改 /home/fyk/my/web3/web3Learning/note/week7/test1/api/internal/config/config.go

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Config struct {
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
	rest.RestConf
}
```

### 2. 确保配置文件（user-api.yaml 或类似文件）中包含 Cache 配置

如果使用配置文件，需要添加相应的 Cache 配置项，例如：

```yaml
Cache:
  - Host: localhost:6379
    Type: redis
```

### 3. 验证修复

修改完成后，重新运行 `go run user.go` 命令检查编译错误是否已解决。