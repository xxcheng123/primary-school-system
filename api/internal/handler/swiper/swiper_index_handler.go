package swiper

import (
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/pkg/httpc"

	"github.com/xxcheng123/primary-school-system/api/internal/logic/swiper"
	"github.com/xxcheng123/primary-school-system/api/internal/svc"
	"github.com/xxcheng123/primary-school-system/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SwiperIndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SwiperIndexReq
		if err := httpx.Parse(r, &req); err != nil {
			httpc.Response(r.Context(), w, errs.ParamParseFailed.Code(), errs.ParamParseFailed.Error(), nil)
			return
		}

		l := swiper.NewSwiperIndexLogic(r.Context(), svcCtx)
		resp, err := l.SwiperIndex(&req)
		if err != nil {
			if s, ok := status.FromError(err); ok {
				httpc.Response(r.Context(), w, s.Code(), s.Message(), nil)
			} else {
				httpc.Response(r.Context(), w, 1000, err.Error(), nil)
			}
		} else {
			httpc.Success(r.Context(), w, resp)
		}
	}
}
