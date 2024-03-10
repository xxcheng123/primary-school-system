package categorylogic

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 15:34:02
 * @LastEditTime: 2024-03-08 15:49:18
 */

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xxcheng123/primary-school-system/ent"
	"github.com/xxcheng123/primary-school-system/ent/category"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListCategoryLogic {
	return &GetListCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListCategoryLogic) GetListCategory(in *pb.GetListCategoryReq) (*pb.ListCategoryResp, error) {
	var index, limit int64 = 0, 10
	query := l.svcCtx.DB.Category.Query()
	if in.Index != nil && *in.Index > 0 {
		index = *in.Index
		query = query.Offset(int(index * limit))
	}
	if in.Limit != nil && *in.Limit > 0 {
		limit = *in.Limit
		query = query.Limit(int(limit))
	}

	if in.ParentId != nil {
		query = query.Where(category.ParentID(*in.ParentId)).Where(category.Status(1))
	}
	count, err := query.Count(l.ctx)
	if err != nil {
		return nil, errs.MySQLInternalError
	}
	list, err := query.All(l.ctx)
	if err != nil {
		return nil, errs.MySQLInternalError
	}
	var respList []*pb.CategoryInfo
	err = copier.CopyWithOption(&respList, list, copier.Option{
		FieldNameMapping: []copier.FieldNameMapping{
			{
				SrcType: ent.Category{},
				DstType: pb.CategoryInfo{},
				Mapping: map[string]string{
					"ID":       "Id",
					"ParentID": "ParentId",
				},
			},
		},
	})
	return &pb.ListCategoryResp{
		Total: int64(count),
		Index: index,
		Limit: limit,
		List:  respList,
	}, nil
}
