// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/a11en4sec/lebron/apps/car/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: AdminHandler(serverCtx),
			},
		},
	)
}
