package userlogic

import (
	"context"
	"fmt"

	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *pb.CreateUserReq) (*pb.CreateUserResp, error) {
	username := in.Username
	password := in.Password
	name := in.Name
	var role int64 = 1
	if in.Role != nil {
		role = *in.Role
	}
	u, err := l.svcCtx.DB.User.Create().
		SetUsername(username).
		SetPassword(password).
		SetName(name).
		SetRole(role).
		Save(l.ctx)
	if err != nil {
		fmt.Println(err)
		l.Logger.Error(err)
		return nil, errs.MySQLInternalError
	}
	return &pb.CreateUserResp{
		Id:       u.ID,
		Username: u.Username,
		Role:     u.Role,
		Name:     u.Name,
	}, nil
}
