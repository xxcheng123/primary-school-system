package categorylogic

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 15:15:35
 * @LastEditTime: 2024-03-08 16:45:58
 */

import (
	"context"

	"github.com/xxcheng123/primary-school-system/ent"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCategoryLogic) CreateCategory(in *pb.CreateCategoryReq) (*pb.CategoryInfo, error) {
	name := in.Name
	var parentId int64 = 0
	if in.ParentId != nil && *in.ParentId != 0 {
		parentId = *in.ParentId
		_, err := l.svcCtx.DB.Category.Get(l.ctx, parentId)
		if ent.IsNotFound(err) {
			return nil, errs.CategoryNotExist
		} else if err != nil {
			l.Logger.Error(err)
			return nil, errs.MySQLInternalError
		}
	}
	c, err := l.svcCtx.DB.Category.Create().
		SetName(name).
		SetParentID(parentId).
		Save(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return nil, errs.MySQLInternalError
	}
	return &pb.CategoryInfo{
		Id:       c.ID,
		Name:     c.Name,
		ParentId: c.ParentID,
		Status:   c.Status,
	}, nil
}
