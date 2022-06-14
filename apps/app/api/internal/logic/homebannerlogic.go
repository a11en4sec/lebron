package logic

import (
	"context"

	"github.com/a11en4sec/lebron/apps/app/api/internal/svc"
	"github.com/a11en4sec/lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeBannerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBannerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBannerLogic {
	return &HomeBannerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBannerLogic) HomeBanner() (resp *types.HomeBannerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
