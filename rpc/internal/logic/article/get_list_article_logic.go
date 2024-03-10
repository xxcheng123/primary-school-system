package articlelogic

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 17:27:39
 * @LastEditTime: 2024-03-08 17:30:58
 */

import (
	"context"

	"github.com/xxcheng123/primary-school-system/ent/article"
	"github.com/xxcheng123/primary-school-system/ent/category"
	"github.com/xxcheng123/primary-school-system/ent/user"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListArticleLogic {
	return &GetListArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListArticleLogic) GetListArticle(in *pb.GetListArticleReq) (*pb.GetListArticleResp, error) {
	var index, limit int64 = 0, 10
	query := l.svcCtx.DB.Article.Query().WithCategory().WithUser()
	if in.Index != nil && *in.Index > 0 {
		index = *in.Index
		query = query.Offset(int(index * limit))
	}
	if in.Limit != nil && *in.Limit > 0 {
		limit = *in.Limit
		query = query.Limit(int(limit))
	}

	if in.UserId != nil {
		query = query.Where(article.HasUserWith(user.ID(*in.UserId)))
	}
	if in.CategoryId != nil {
		ids := []int64{*in.CategoryId}
		if in.IncludeSubCategory != nil && *in.IncludeSubCategory {
			for i := 0; i < len(ids); {
				subList, err := l.svcCtx.DB.Category.Query().
					Where(category.ParentIDIn(ids[i:]...)).
					All(l.ctx)
				if err == nil && len(subList) > 0 {
					for j := 0; j < len(subList); j++ {
						ids = append(ids, subList[j].ID)
					}
				}
				i = len(ids)
			}
		}
		query = query.Where(article.HasCategoryWith(category.IDIn(ids...)))
	}
	count, err := query.Count(l.ctx)
	if err != nil {
		return nil, errs.MySQLInternalError
	}
	list, err := query.All(l.ctx)
	if err != nil {
		return nil, errs.MySQLInternalError
	}
	var respList []*pb.ArticleInfo
	for _, info := range list {
		respList = append(respList, &pb.ArticleInfo{
			Id:         info.ID,
			Title:      info.Title,
			Content:    info.Content,
			Status:     info.Status,
			UserId:     info.Edges.User.ID,
			CategoryId: info.Edges.Category.ID,
			Img:        info.Img,
		})
	}
	return &pb.GetListArticleResp{
		Total: int64(count),
		List:  respList,
		Index: index,
		Limit: limit,
	}, nil
}
