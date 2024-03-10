package swiper

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/xxcheng123/primary-school-system/api/internal/svc"
	"github.com/xxcheng123/primary-school-system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwiperIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwiperIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwiperIndexLogic {
	return &SwiperIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwiperIndexLogic) SwiperIndex(req *types.SwiperIndexReq) (resp *types.SwiperIndexResp, err error) {
	p := &pb.ListSwiperReq{
		Index: &req.Index,
		Limit: nil,
	}
	if req.Size > 0 {
		p.Limit = &req.Size
	}
	rpcResult, err := l.svcCtx.SchoolRpc.Swiper.ListSwiper(l.ctx, p)
	if err != nil {
		return nil, err
	}
	resp = &types.SwiperIndexResp{}
	_ = copier.Copy(resp, rpcResult)
	resp.Size = rpcResult.Limit
	return
}
