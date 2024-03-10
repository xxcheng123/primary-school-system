package articlelogic

/*
 * @Author: xxcheng
 * @Email: developer@xxcheng.cn
 * @Date: 2024-03-08 15:43:51
 * @LastEditTime: 2024-03-08 15:45:13
 */

import (
	"context"

	"github.com/xxcheng123/primary-school-system/pkg/errs"
	"github.com/xxcheng123/primary-school-system/rpc/internal/svc"
	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateArticleLogic) CreateArticle(in *pb.CreateArticleReq) (*pb.CreateArticleResp, error) {
	uid := in.UserId
	cid := in.CategoryId
	title := in.Title
	content := in.Content
	var status int64 = 1
	if in.Status != nil {
		status = *in.Status
	}
	img := "https://zip.xxcheng.cn/lunwen/news.jpeg"
	if in.Img != nil {
		img = *in.Img
	}
	a, err := l.svcCtx.DB.Article.Create().
		SetUserID(uid).
		SetCategoryID(cid).
		SetTitle(title).
		SetContent(content).
		SetStatus(status).
		SetImg(img).
		Save(l.ctx)
	if err != nil {
		l.Logger.Error(err)
		return nil, errs.MySQLInternalError
	}
	return &pb.CreateArticleResp{
		Id: a.ID,
	}, nil
}
