package userlogic

import (
	"context"

	"github.com/xxcheng123/primary-school-system/ent"
	"github.com/xxcheng123/primary-school-system/ent/user"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginUserLogic) LoginUser(in *pb.LoginUserReq) (*pb.InfoUserResp, error) {
	username := in.Username
	password := in.Password
	u, err := l.svcCtx.DB.User.Query().Where(user.Username(username)).
		Only(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.UserNotExist
		}
		return nil, errs.MySQLInternalError
	}
	if u.Password != password {
		return nil, errs.PasswordIncorrect
	}
	if u.Status < 0 {
		return nil, errs.UserBan
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
