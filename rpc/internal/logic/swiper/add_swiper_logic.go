package swiperlogic

import (
	"context"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/pkg/fcp"

	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSwiperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSwiperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSwiperLogic {
	return &AddSwiperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSwiperLogic) AddSwiper(in *pb.AddSwiperReq) (*pb.InfoSwiper, error) {
	query := l.svcCtx.DB.Swiper.Create().SetTitle(in.Title).SetURL(in.Url)
	if in.Img != nil {
		query = query.SetImg(*in.Img)
	}
	if in.Order != nil {
		query = query.SetOrder(*in.Order)
	}
	if in.Status != nil {
		query = query.SetStatus(*in.Status)
	}
	s, err := query.Save(l.ctx)
	if err != nil {
		return nil, errs.MySQLInternalError
	}
	return fcp.ConvSwiper(s)
}
