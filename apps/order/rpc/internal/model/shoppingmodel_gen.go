// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	shoppingFieldNames          = builder.RawFieldNames(&Shopping{})
	shoppingRows                = strings.Join(shoppingFieldNames, ",")
	shoppingRowsExpectAutoSet   = strings.Join(stringx.Remove(shoppingFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	shoppingRowsWithPlaceHolder = strings.Join(stringx.Remove(shoppingFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheOrdersShoppingIdPrefix = "cache:orders:shopping:id:"
)

type (
	shoppingModel interface {
		Insert(ctx context.Context, data *Shopping) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Shopping, error)
		Update(ctx context.Context, newData *Shopping) error
		Delete(ctx context.Context, id int64) error
	}

	defaultShoppingModel struct {
		sqlc.CachedConn
		table string
	}

	Shopping struct {
		Id               int64     `db:"id"`                // 收货信息表id
		Orderid          string    `db:"orderid"`           // 订单id
		Userid           int64     `db:"userid"`            // 用户id
		ReceiverName     string    `db:"receiver_name"`     // 收货姓名
		ReceiverPhone    string    `db:"receiver_phone"`    // 收货固定电话
		ReceiverMobile   string    `db:"receiver_mobile"`   // 收货移动电话
		ReceiverProvince string    `db:"receiver_province"` // 省份
		ReceiverCity     string    `db:"receiver_city"`     // 城市
		ReceiverDistrict string    `db:"receiver_district"` // 区/县
		ReceiverAddress  string    `db:"receiver_address"`  // 详细地址
		CreateTime       time.Time `db:"create_time"`       // 创建时间
		UpdateTime       time.Time `db:"update_time"`       // 更新时间
	}
)

func newShoppingModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultShoppingModel {
	return &defaultShoppingModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`shopping`",
	}
}

func (m *defaultShoppingModel) Delete(ctx context.Context, id int64) error {
	ordersShoppingIdKey := fmt.Sprintf("%s%v", cacheOrdersShoppingIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, ordersShoppingIdKey)
	return err
}

func (m *defaultShoppingModel) FindOne(ctx context.Context, id int64) (*Shopping, error) {
	ordersShoppingIdKey := fmt.Sprintf("%s%v", cacheOrdersShoppingIdPrefix, id)
	var resp Shopping
	err := m.QueryRowCtx(ctx, &resp, ordersShoppingIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shoppingRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultShoppingModel) Insert(ctx context.Context, data *Shopping) (sql.Result, error) {
	ordersShoppingIdKey := fmt.Sprintf("%s%v", cacheOrdersShoppingIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, shoppingRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Orderid, data.Userid, data.ReceiverName, data.ReceiverPhone, data.ReceiverMobile, data.ReceiverProvince, data.ReceiverCity, data.ReceiverDistrict, data.ReceiverAddress)
	}, ordersShoppingIdKey)
	return ret, err
}

func (m *defaultShoppingModel) Update(ctx context.Context, data *Shopping) error {
	ordersShoppingIdKey := fmt.Sprintf("%s%v", cacheOrdersShoppingIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, shoppingRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Orderid, data.Userid, data.ReceiverName, data.ReceiverPhone, data.ReceiverMobile, data.ReceiverProvince, data.ReceiverCity, data.ReceiverDistrict, data.ReceiverAddress, data.Id)
	}, ordersShoppingIdKey)
	return err
}

func (m *defaultShoppingModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheOrdersShoppingIdPrefix, primary)
}

func (m *defaultShoppingModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", shoppingRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultShoppingModel) tableName() string {
	return m.table
}
