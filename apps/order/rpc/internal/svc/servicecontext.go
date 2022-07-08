package svc

import (
	"github.com/a11en4sec/lebron/apps/order/rpc/internal/config"
	"github.com/a11en4sec/lebron/apps/order/rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrdersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrdersModel(conn, c.CacheRedis),
	}
}
