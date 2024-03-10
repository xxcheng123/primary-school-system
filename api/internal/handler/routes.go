// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	swiper "github.com/xxcheng123/primary-school-system/api/internal/handler/swiper"
	"github.com/xxcheng123/primary-school-system/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: PingHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/index",
				Handler: swiper.SwiperIndexHandler(serverCtx),
			},
		},
		rest.WithPrefix("/swiper"),
	)
}
