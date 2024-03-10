package tpl

import (
	"github.com/xxcheng123/primary-school-system/api/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"html/template"
	"net/http"
	"time"
)

/**
 * @Author: root
 * @Description:
 * @File:  router
 * @Date: 3/8/24 10:34 PM
 */

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext, tmplDir string) {
	tpl, err := template.New("school").ParseGlob(tmplDir)
	if err != nil {
		panic(err)
	}
	server.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/index.html",
			Handler: IndexHandler(tpl, serverCtx),
		},
		{
			Method:  http.MethodGet,
			Path:    "/test.html",
			Handler: DebugHandler(tpl, serverCtx),
		},
	})
	go func() {
		for {
			time.Sleep(time.Second)
			tpl2, err := template.New("school").ParseGlob(tmplDir)
			if err != nil {
				continue
			}
			*tpl = *tpl2
		}
	}()
}
