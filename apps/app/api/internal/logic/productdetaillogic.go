package logic

import (
	"context"

	"github.com/a11en4sec/lebron/apps/app/api/internal/svc"
	"github.com/a11en4sec/lebron/apps/app/api/internal/types"
	"github.com/a11en4sec/lebron/apps/product/rpc/product"
	"github.com/a11en4sec/lebron/apps/reply/rpc/reply"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type ProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDetailLogic) ProductDetail(req *types.ProductDetailRequest) (resp *types.ProductDetailResponse, err error) {
	// todo: add your logic here and delete this line
	var (
		p *product.ProductItem
		cs *reply.CommentsResponse
	)

	if err := mr.Finish(func() error {
		var err error
		if p, err = l.svcCtx.ProductRPC.Product(l.ctx,&product.ProductItemRequest{ProductId: req.ProductID});err != nil {
			return err
		}
		return nil 
	}, func() error {
		var err error
		if cs, err = l.svcCtx.ReplyRPC.Comments(l.ctx, &reply.CommentsRequest{TargetId: req.ProductID});err != nil {
			logx.Errorf("get commands error:%v",err)
		}
		return nil 
	}); err != nil {
		return nil, err
	}

	var comments []*types.Comment

	for _, c := range cs.Comments{
		comments = append(comments, &types.Comment{
			ID:         c.Id,
			Content:    c.Content,

		})
	}

	return &types.ProductDetailResponse{
		Product:  &types.Product{
			ID:          p.ProductId,
			Name:        p.Name,
			// Images:      []string{},
			// Description: "",
			// Price:       0,
			// Stock:       0,
			// Category:    "",
			// Status:      0,
			// CreateTime:  0,
			// UpdateTime:  0,
		},
		Comments: comments,
	},nil
}
