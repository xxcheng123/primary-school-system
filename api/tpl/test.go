package tpl

import (
	"github.com/xxcheng123/primary-school-system/api/internal/svc"
	"html/template"
	"net/http"
)

/**
 * @Author: root
 * @Description:
 * @File:  test
 * @Date: 3/10/24 3:52 PM
 */

func DebugHandler(tpl *template.Template, svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Info struct {
			Name string `json:"name"`
		}
		list := []*Info{
			{Name: "xxw"},
			{Name: "wdew"},
			{Name: "wdewf"},
		}
		tpl.ExecuteTemplate(w, "test.html", list)
	}
}
