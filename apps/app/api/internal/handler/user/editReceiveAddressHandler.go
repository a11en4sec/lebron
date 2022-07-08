package user

import (
	"net/http"

	"github.com/a11en4sec/lebron/apps/app/api/internal/logic/user"
	"github.com/a11en4sec/lebron/apps/app/api/internal/svc"
	"github.com/a11en4sec/lebron/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EditReceiveAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReceiveAddressEditReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewEditReceiveAddressLogic(r.Context(), svcCtx)
		resp, err := l.EditReceiveAddress(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
