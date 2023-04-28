package main

import (
	"github.com/labstack/echo/v4"
	"magic_box/server/handler/tools"
	"magic_box/server/handler/user"
	"magic_box/server/service"
)

func initRoute(e *echo.Echo, svc *service.Service) {
	var (
		toolHdl = tools.NewTool(svc)
		userHdl = user.NewUser(svc)
	)

	var (
		userGroup = e.Group("/user")
		toolGroup = e.Group("/tools")
	)

	userGroup.POST("/login", userHdl.Login)
	userGroup.POST("/refresh_access_token", userHdl.RefreshAccessToken)

	toolGroup.POST("/base64", toolHdl.Base64)
	toolGroup.POST("/json", toolHdl.JsonPretty)
	toolGroup.POST("/url", toolHdl.Url)

}
