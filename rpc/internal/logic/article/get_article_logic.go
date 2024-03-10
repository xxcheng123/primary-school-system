package articlelogic

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 15:47:18
 * @LastEditTime: 2024-03-08 17:21:09
 */

import (
	"context"

	"github.com/xxcheng123/primary-school-system/ent"
	"github.com/xxcheng123/primary-school-system/ent/article"
	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetArticleLogic) GetArticle(in *pb.GetInfoArticleReq) (*pb.ArticleInfo, error) {
	id := in.Id
	c, err := l.svcCtx.DB.Article.Query().WithCategory().WithUser().Where(article.ID(id)).First(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errs.NotFound
		}
		return nil, errs.MySQLInternalError
	}
	return &pb.ArticleInfo{
		Id:         c.ID,
		UserId:     c.Edges.User.ID,
		CategoryId: c.Edges.Category.ID,
		Title:      c.Title,
		Content:    c.Content,
		Status:     c.Status,
		Img:        c.Img,
	}, nil
}
