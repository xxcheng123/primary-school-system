package categorylogic

import (
	"context"

	"github.com/xxcheng123/primary-school-system/ent"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCategoryLogic {
	return &GetCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCategoryLogic) GetCategory(in *pb.GetCategoryReq) (*pb.CategoryInfo, error) {
	id := in.Id
	c, err := l.svcCtx.DB.Category.Get(l.ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.NotFound
		}
		return nil, errs.MySQLInternalError
	}
	return &pb.CategoryInfo{
		Id:       c.ID,
		Name:     c.Name,
		ParentId: c.ParentID,
		Status:   c.Status,
	}, nil
}
