package userlogic

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 15:00:29
 * @LastEditTime: 2024-03-08 16:22:42
 */

import (
	"context"

	"github.com/xxcheng123/primary-school-system/ent"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInfoUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoUserLogic {
	return &GetInfoUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInfoUserLogic) GetInfoUser(in *pb.GetInfoUserReq) (*pb.InfoUserResp, error) {
	id := in.Id
	u, err := l.svcCtx.DB.User.Get(l.ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.UserNotExist
		}
		return nil, errs.MySQLInternalError
	}
	return &pb.InfoUserResp{
		Info: &pb.UserInfo{
			Id:        u.ID,
			Username:  u.Username,
			Role:      u.Role,
			Status:    u.Status,
			UpdatedAt: u.UpdateTime.UnixMilli(),
			Name:      u.Name,
		},
		Code: 200,
	}, nil
}
