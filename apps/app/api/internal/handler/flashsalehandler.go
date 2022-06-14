package handler

import (
	"net/http"

	"github.com/a11en4sec/lebron/apps/app/api/internal/logic"
	"github.com/a11en4sec/lebron/apps/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FlashSaleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFlashSaleLogic(r.Context(), svcCtx)
		resp, err := l.FlashSale()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
