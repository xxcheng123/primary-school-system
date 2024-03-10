package main

import (
	"flag"
	"fmt"
	"github.com/xxcheng123/primary-school-system/api/internal/config"
	"github.com/xxcheng123/primary-school-system/api/internal/handler"
	"github.com/xxcheng123/primary-school-system/api/internal/svc"
	"github.com/xxcheng123/primary-school-system/api/tpl"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/school.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	tpl.RegisterHandlers(server, ctx, "template/*.html")

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
