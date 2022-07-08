package logic

import (
	"context"
	"strings"

	"github.com/a11en4sec/lebron/apps/product/rpc/internal/model"
	"github.com/a11en4sec/lebron/apps/product/rpc/internal/svc"
	"github.com/a11en4sec/lebron/apps/product/rpc/product"
	"github.com/zeromicro/go-zero/core/mr"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {
	//数ProductIds为逗号分隔的多个商品id，在这里我们使用go-zero提供的mapreduce来并行的根据商品id获取商品详情
	products := make(map[int64]*product.ProductItem)

	pdis := strings.Split(in.ProductIds, ",")
	ps, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, pid := range pdis {
			source <- pid
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		pid := item.(int64)
		p, err := l.svcCtx.ProductModel.FindOne(l.ctx, pid)
		if err != nil {
			cancel(err)
			return
		}
		writer.Write(p)
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var r []*model.Product
		for p := range pipe {
			r = append(r, p.(*model.Product))
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	for _, p := range ps.([]*model.Product) {
		products[p.Id] = &product.ProductItem{
			ProductId: p.Id,
			Name:      p.Name,
		}
	}
	return &product.ProductResponse{Products: products}, nil
}
