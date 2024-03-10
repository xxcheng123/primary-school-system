package swiperlogic

import (
	"context"
	"github.com/xxcheng123/primary-school-system/ent"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/pkg/fcp"

	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSwiperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSwiperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSwiperLogic {
	return &UpdateSwiperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSwiper 更新轮播图
func (l *UpdateSwiperLogic) UpdateSwiper(in *pb.UpdateSwiperReq) (*pb.InfoSwiper, error) {
	query := l.svcCtx.DB.Swiper.UpdateOneID(in.Id)
	if in.Title != nil {
		query = query.SetTitle(*in.Title)
	}
	if in.Url != nil {
		query = query.SetURL(*in.Url)
	}
	if in.Img != nil {
		query = query.SetImg(*in.Img)
	}
	if in.Order != nil {
		query = query.SetOrder(*in.Order)
	}
	if in.Status != nil {
		query = query.SetStatus(*in.Status)
	}
	err := query.Exec(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.NotFound
		}
		return nil, errs.OperateFailure
	}
	s, err := l.svcCtx.DB.Swiper.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errs.OperateFailure
	}
	return fcp.ConvSwiper(s)
}
