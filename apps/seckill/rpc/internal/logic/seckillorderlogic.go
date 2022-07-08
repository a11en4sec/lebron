package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/a11en4sec/lebron/apps/product/rpc/product"
	"github.com/a11en4sec/lebron/apps/seckill/rpc/internal/svc"
	"github.com/a11en4sec/lebron/apps/seckill/rpc/seckill"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/limit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	limitPeriod       = 10
	limitQuota        = 1
	seckillUserPrefix = "seckill#u#"
	localCacheExpire  = time.Second * 60
)

type SeckillOrderLogic struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	limiter    *limit.PeriodLimit
	localCache *collection.Cache
	logx.Logger
}

type KafkaData struct {
	Uid int64 `json:"uid"`
	Pid int64 `json:"pid"`
}

func NewSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillOrderLogic {
	localCache, err := collection.NewCache(localCacheExpire)
	if err != nil {
		panic(err)
	}

	return &SeckillOrderLogic{
		ctx:        ctx,
		svcCtx:     svcCtx,
		Logger:     logx.WithContext(ctx),
		localCache: localCache,
		limiter:    limit.NewPeriodLimit(limitPeriod, limitQuota, svcCtx.BizRedis, seckillUserPrefix),
	}
}

func (l *SeckillOrderLogic) SeckillOrder(in *seckill.SeckillOrderRequest) (*seckill.SeckillOrderResponse, error) {

	// 1 periodlimit限流器,,可能返回1(in limit),2(hit limit),3(over limit)
	code, _ := l.limiter.Take(strconv.FormatInt(in.UserId, 10))
	if code == limit.OverQuota {
		return nil, status.Errorf(codes.OutOfRange, "Number of requests exceeded the limit")
	}

	p, err := l.svcCtx.ProductRPC.Product(l.ctx, &product.ProductItemRequest{ProductId: in.ProductId})
	if err != nil {
		return nil, err
	}

	// fmt.Printf("[p]:%v\n", p)
	// fmt.Printf("[p.stock]:%d\n",p.Stock)

	// 没有库存,直接返回
	if p.Stock <= 0 {
		return nil, status.Errorf(codes.OutOfRange, "Insufficient stock")
	}

	kd, err := json.Marshal(&KafkaData{Uid: in.UserId, Pid: in.ProductId})

	if err != nil {
		return nil, err
	}

	fmt.Println("kafka name",l.svcCtx.KafkaPusher.Name())
	
	if err := l.svcCtx.KafkaPusher.Push(string(kd)); err != nil {
		fmt.Printf("kafka error:%v\n",err.Error())
		return nil, err
	}

	return &seckill.SeckillOrderResponse{}, nil
}
