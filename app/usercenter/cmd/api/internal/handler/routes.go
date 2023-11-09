// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "automatix/app/usercenter/cmd/api/internal/handler/user"
	"automatix/app/usercenter/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/fastLogin",
				Handler: user.FastLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: user.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/wxMiniAuth",
				Handler: user.WxMiniAuthHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/usercenter/v1"),
	)
}
