package swiperlogic

import (
	"context"
	"github.com/xxcheng123/primary-school-system/ent"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	exec2 "github.com/xxcheng123/primary-school-system/pkg/exec"

	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperateSwiperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperateSwiperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperateSwiperLogic {
	return &OperateSwiperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OperateSwiper 简单编辑轮播图
func (l *OperateSwiperLogic) OperateSwiper(in *pb.OperateSwiperReq) (*pb.OperateSwiperResp, error) {
	const OpTypUp, OpTypeDown, OpTypeDel = "up", "down", "del"
	opType := in.OpType
	id := in.Id
	query := l.svcCtx.DB.Swiper
	var exec exec2.ExecIF
	switch opType {
	case OpTypUp:
		exec = query.UpdateOneID(id).SetStatus(1)
	case OpTypeDown:
		exec = query.UpdateOneID(id).SetStatus(2)
	case OpTypeDel:
		exec = query.DeleteOneID(id)
	default:
		return nil, errs.NotSupportedOperate
	}
	err := exec.Exec(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.NotFound
		}
		return nil, errs.OperateFailure
	}
	return &pb.OperateSwiperResp{
		Code: 200,
		Id:   in.Id,
	}, nil
}
