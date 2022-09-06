package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/imgbjs/imgbjs/api/internal/config"
	"github.com/imgbjs/imgbjs/api/internal/handler"
	"github.com/imgbjs/imgbjs/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/imgbjs-api.yaml", "the config file")

var Config config.Config

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	Config = c

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/js/:file",
				Handler: http.StripPrefix("/js/", http.FileServer(http.Dir("view/js"))).ServeHTTP,
			},
			{
				Method:  http.MethodGet,
				Path:    "/style/:file",
				Handler: http.StripPrefix("/style/", http.FileServer(http.Dir("view/style"))).ServeHTTP,
			},
			{
				Method:  http.MethodGet,
				Path:    "/meta/:file",
				Handler: http.StripPrefix("/meta/", http.FileServer(http.Dir("view/meta"))).ServeHTTP,
			},
		},
	)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
