// Code generated by goctl. DO NOT EDIT.
// Source: school.proto

package category

import (
	"context"

	"github.com/xxcheng123/primary-school-system/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddSwiperReq       = pb.AddSwiperReq
	ArticleInfo        = pb.ArticleInfo
	CategoryInfo       = pb.CategoryInfo
	CreateArticleReq   = pb.CreateArticleReq
	CreateArticleResp  = pb.CreateArticleResp
	CreateCategoryReq  = pb.CreateCategoryReq
	CreateUserReq      = pb.CreateUserReq
	CreateUserResp     = pb.CreateUserResp
	GetCategoryReq     = pb.GetCategoryReq
	GetInfoArticleReq  = pb.GetInfoArticleReq
	GetInfoUserReq     = pb.GetInfoUserReq
	GetListArticleReq  = pb.GetListArticleReq
	GetListArticleResp = pb.GetListArticleResp
	GetListCategoryReq = pb.GetListCategoryReq
	InfoSwiper         = pb.InfoSwiper
	InfoUserResp       = pb.InfoUserResp
	ListCategoryResp   = pb.ListCategoryResp
	ListSwiperReq      = pb.ListSwiperReq
	ListSwiperResp     = pb.ListSwiperResp
	LoginUserReq       = pb.LoginUserReq
	OperateSwiperReq   = pb.OperateSwiperReq
	OperateSwiperResp  = pb.OperateSwiperResp
	Request            = pb.Request
	Response           = pb.Response
	UpdateSwiperReq    = pb.UpdateSwiperReq
	UserInfo           = pb.UserInfo

	Category interface {
		CreateCategory(ctx context.Context, in *CreateCategoryReq, opts ...grpc.CallOption) (*CategoryInfo, error)
		GetCategory(ctx context.Context, in *GetCategoryReq, opts ...grpc.CallOption) (*CategoryInfo, error)
		GetListCategory(ctx context.Context, in *GetListCategoryReq, opts ...grpc.CallOption) (*ListCategoryResp, error)
	}

	defaultCategory struct {
		cli zrpc.Client
	}
)

func NewCategory(cli zrpc.Client) Category {
	return &defaultCategory{
		cli: cli,
	}
}

func (m *defaultCategory) CreateCategory(ctx context.Context, in *CreateCategoryReq, opts ...grpc.CallOption) (*CategoryInfo, error) {
	client := pb.NewCategoryClient(m.cli.Conn())
	return client.CreateCategory(ctx, in, opts...)
}

func (m *defaultCategory) GetCategory(ctx context.Context, in *GetCategoryReq, opts ...grpc.CallOption) (*CategoryInfo, error) {
	client := pb.NewCategoryClient(m.cli.Conn())
	return client.GetCategory(ctx, in, opts...)
}

func (m *defaultCategory) GetListCategory(ctx context.Context, in *GetListCategoryReq, opts ...grpc.CallOption) (*ListCategoryResp, error) {
	client := pb.NewCategoryClient(m.cli.Conn())
	return client.GetListCategory(ctx, in, opts...)
}