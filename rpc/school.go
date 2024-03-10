package main

import (
	"flag"
	"fmt"

	"github.com/xxcheng123/primary-school-system/rpc/internal/config"
	articleServer "github.com/xxcheng123/primary-school-system/rpc/internal/server/article"
	baseServer "github.com/xxcheng123/primary-school-system/rpc/internal/server/base"
	categoryServer "github.com/xxcheng123/primary-school-system/rpc/internal/server/category"
	swiperServer "github.com/xxcheng123/primary-school-system/rpc/internal/server/swiper"
	userServer "github.com/xxcheng123/primary-school-system/rpc/internal/server/user"
	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/school.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterBaseServer(grpcServer, baseServer.NewBaseServer(ctx))
		pb.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		pb.RegisterCategoryServer(grpcServer, categoryServer.NewCategoryServer(ctx))
		pb.RegisterArticleServer(grpcServer, articleServer.NewArticleServer(ctx))
		pb.RegisterSwiperServer(grpcServer, swiperServer.NewSwiperServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
