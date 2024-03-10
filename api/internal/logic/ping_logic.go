package logic

import (
	"context"

	"github.com/xxcheng123/primary-school-system/api/internal/svc"
	"github.com/xxcheng123/primary-school-system/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.Ping) (resp *types.Pong, err error) {
	return &types.Pong{
		Pong: "pong",
	}, nil
}
