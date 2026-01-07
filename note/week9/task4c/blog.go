package main

import (
	"flag"
	"fmt"

	"task4c/internal/config"
	"task4c/internal/handler"
	"task4c/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/blog.yaml", "the config file")

// main函数是程序的入口点，负责初始化配置、创建服务器并启动服务
func main() {
	// 解析命令行参数，用于读取配置文件路径等参数
	flag.Parse()

	// 声明配置变量
	var c config.Config
	// 加载配置文件内容到配置变量中，使用MustLoad确保加载失败时程序会终止
	conf.MustLoad(*configFile, &c)

	// 创建RESTful服务器实例，使用配置中的RestConf参数
	server := rest.MustNewServer(c.RestConf)
	// 确保程序退出前停止服务器，释放资源
	defer server.Stop()

	// 创建服务上下文，包含数据库连接、缓存等共享资源
	ctx := svc.NewServiceContext(c)
	// 注册HTTP处理器到服务器，将请求路由到对应的处理函数
	handler.RegisterHandlers(server, ctx)

	// 打印服务器启动信息，显示监听的主机和端口
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// 启动服务器，开始监听HTTP请求
	server.Start()
}
