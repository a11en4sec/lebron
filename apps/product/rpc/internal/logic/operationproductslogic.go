package logic

import (
	"context"

	"github.com/a11en4sec/lebron/apps/product/rpc/internal/svc"
	"github.com/a11en4sec/lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationProductsLogic struct {
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	productListLogic *ProductListLogic
	logx.Logger
}

func NewOperationProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationProductsLogic {
	return &OperationProductsLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		productListLogic: NewProductListLogic(ctx, svcCtx),
	}
}

const (
	validStatus          = 1
	operationProductsKey = "operation#products"
)

func (l *OperationProductsLogic) OperationProducts(in *product.OperationProductsRequest) (*product.OperationProductsResponse, error) {
	// 1 从缓存查
	opProducts, ok := l.svcCtx.LocalCache.Get(operationProductsKey)

	if ok {
		return &product.OperationProductsResponse{
			Products: opProducts.([]*product.ProductItem),
		}, nil
	}

	// 2 缓存查不到，查数据库
	pos, err := l.svcCtx.OperationModel.OperationProducts(l.ctx,validStatus)
	if err != nil {
		return nil, err
	}

	var pids []int64
	for _, p := range pos {
		pids = append(pids, p.ProductId)
	}

	// 3 按pids批量查产品信息
	products, err := l.productListLogic.productsByIds(l.ctx, pids)
	if err != nil {
		return nil, err
	}

	var pItems []*product.ProductItem
	for _, p := range products {
		pItems = append(pItems, &product.ProductItem{
			ProductId: p.Id,
			Name: p.Name,
		})
	}
	l.svcCtx.LocalCache.Set(operationProductsKey,pItems)

	return &product.OperationProductsResponse{}, nil
}
