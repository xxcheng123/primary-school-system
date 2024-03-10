package svc

import (
	"github.com/xxcheng123/primary-school-system/api/internal/config"
	"github.com/xxcheng123/primary-school-system/rpc/client/article"
	"github.com/xxcheng123/primary-school-system/rpc/client/category"
	"github.com/xxcheng123/primary-school-system/rpc/client/swiper"
	"github.com/xxcheng123/primary-school-system/rpc/client/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	SchoolRpc *struct {
		article.Article
		user.User
		swiper.Swiper
		category.Category
	}
}

func NewServiceContext(c config.Config) *ServiceContext {
	schoolRpc := &struct {
		article.Article
		user.User
		swiper.Swiper
		category.Category
	}{
		article.NewArticle(zrpc.MustNewClient(c.SchoolRpcConf)),
		user.NewUser(zrpc.MustNewClient(c.SchoolRpcConf)),
		swiper.NewSwiper(zrpc.MustNewClient(c.SchoolRpcConf)),
		category.NewCategory(zrpc.MustNewClient(c.SchoolRpcConf)),
	}
	return &ServiceContext{
		Config:    c,
		SchoolRpc: schoolRpc,
	}
}
