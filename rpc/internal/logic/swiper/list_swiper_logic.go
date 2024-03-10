package swiperlogic

import (
	"context"
	"github.com/xxcheng123/primary-school-system/ent"
	"github.com/xxcheng123/primary-school-system/ent/swiper"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/pkg/fcp"

	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSwiperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListSwiperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSwiperLogic {
	return &ListSwiperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListSwiper 轮播图列表
func (l *ListSwiperLogic) ListSwiper(in *pb.ListSwiperReq) (*pb.ListSwiperResp, error) {
	var index, limit int64 = 0, 10
	query := l.svcCtx.DB.Swiper.Query()
	if in.Index != nil && *in.Index > 0 {
		index = *in.Index
		query = query.Offset(int(index * limit))
	}
	if in.Limit != nil && *in.Limit > 0 {
		limit = *in.Limit
		query = query.Limit(int(limit))
	}

	count, err := query.Count(l.ctx)
	if err != nil {
		return nil, errs.MySQLInternalError
	}
	list, err := query.
		Where(swiper.Status(1)).
		Order(ent.Desc(swiper.FieldOrder, swiper.FieldCreateTime)).
		All(l.ctx)
	if err != nil {
		return nil, errs.MySQLInternalError
	}
	var respList []*pb.InfoSwiper
	for _, s := range list {
		item, _ := fcp.ConvSwiper(s)
		respList = append(respList, item)
	}
	return &pb.ListSwiperResp{
		Total: int64(count),
		Index: index,
		Limit: limit,
		List:  respList,
	}, nil
}
