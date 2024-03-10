package tpl

import (
	"encoding/json"
	"fmt"
	"github.com/xxcheng123/primary-school-system/api/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/client/category"
	"github.com/xxcheng123/primary-school-system/types/tmpl"
	"html/template"
	"net/http"
)

/**
 * @Author: root
 * @Description:
 * @File:  main
 * @Date: 3/8/24 10:30 PM
 */

func IndexHandler(tpl *template.Template, svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		pageData := tmpl.Index{
			Base: tmpl.Base{
				Name:  "黄师小学在线网站",
				Title: "当前页面标题",
			},
		}
		// 轮播图数据处理
		{
			rpcResult, err := svcCtx.SchoolRpc.Swiper.ListSwiper(ctx, nil)
			if err == nil {
				pageData.Swiper = rpcResult
			}
		}
		//分类
		{
			rpcResult, _ := svcCtx.SchoolRpc.Category.GetListCategory(ctx, &category.GetListCategoryReq{
				Limit: nil,
			})

			pageData.Category = &tmpl.Category{
				ListCategoryResp: rpcResult,
			}
			root, err := tmpl.BuildCategoryTree(rpcResult.List)
			if err == nil {
				pageData.Category.List = root.Child
			}
		}
		bs, err := json.Marshal(pageData)
		m := map[string]any{}
		_ = json.Unmarshal(bs, &m)
		if err == nil {
			pageData.JSON = m
		} else {
			pageData.JSON = "{}"
		}
		err = tpl.ExecuteTemplate(w, "index.html", pageData)
		if err != nil {
			fmt.Println(err)
		}
	}
}
