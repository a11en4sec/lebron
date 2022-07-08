package service

import (
	"context"
	"encoding/json"

	"github.com/a11en4sec/lebron/apps/order/rpc/order"
	"github.com/a11en4sec/lebron/apps/product/rpc/product"
	"github.com/a11en4sec/lebron/apps/seckill/rmq/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Service struct {
	c          config.Config
	ProductRPC product.Product
	OrderRPC   order.Order
}

type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

func NewService(c config.Config) *Service {
	return &Service{
		c:          c,
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
	}
}

// 消费
func (s *Service) Consume(_ string, value string) error {

	logx.Infof("Consume value: %s\n", value)

	var data KafkaData

	// 反序列化value 到 data结构体中
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}

	// 查询产品是否有库存
	p, err := s.ProductRPC.Product(context.Background(), &product.ProductItemRequest{ProductId: data.Pid})
	if err != nil {
		return err
	}
	if p.Stock <= 0 {
		return nil
	}

	// 创建订单
	_, err = s.OrderRPC.CreateOrder(context.Background(), &order.CreateOrderRequest{Uid: data.Uid, Pid: data.Pid})
	if err != nil {
		logx.Errorf("CreateOrder uid: %d pid: %d error: %v", data.Uid, data.Pid, err)
		return err
	}

	// 扣减库存
	_, err = s.ProductRPC.UpdateProductStock(context.Background(), &product.UpdateProductStockRequest{ProductId: data.Pid, Num: 1})
	if err != nil {
		logx.Errorf("UpdateProductStock uid: %d pid: %d error: %v", data.Uid, data.Pid, err)
		return err
	}

	// TODO notify user of successful order placement

	return nil

}
