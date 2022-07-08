package svc

import (
	"github.com/a11en4sec/lebron/apps/user/rpc/model"
	"github.com/a11en4sec/lebron/apps/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	//add dependency on user model
	UserModel model.UserModel
	//add dependency on user model
	UserReceiveAddressModel model.UserReceiveAddressModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                  c,
		UserModel:               model.NewUserModel(sqlConn, c.CacheRedis),
		UserReceiveAddressModel: model.NewUserReceiveAddressModel(sqlConn, c.CacheRedis),
	}
}
