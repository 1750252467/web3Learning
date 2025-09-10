package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Config struct {
	Name string
	Host string `json:",default=0.0.0.0"`
	Port int
}

var f = flag.String("f", "config.yaml", "config file")

func main() {
	var restConf rest.RestConf
	conf.MustLoad(*f, &restConf)
	s, err := rest.NewServer(restConf)
	if err != nil {
		log.Fatal(err)
		return
	}
	s.AddRoute(rest.Route{
		Path:   "/hello",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			httpx.OkJson(w, "hello word!")
		},
	})
	defer s.Stop()

	s.Start() // 启动服务
}
