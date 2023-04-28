package user

import (
	"github.com/labstack/echo/v4"
	"magic_box/pkg/util"
	"magic_box/server/handler/params"
)

func (h User) Login(c echo.Context) error {
	var (
		apiReq params.LoginReq
		apiRes params.LoginResp
		err    error
	)
	if err = c.Bind(&apiReq); err != nil {
		return err
	}

	apiRes.AccessToken = util.UUIDStr()
	apiRes.ExpireIn = 3600

	return h.RespJson(c, apiRes, err)
}

func (h User) RefreshAccessToken(c echo.Context) error {
	var (
		apiReq params.RefreshAccessTokenReq
		apiRes params.RefreshAccessTokenResp
		err    error
	)
	if err = c.Bind(&apiReq); err != nil {
		return err
	}
	apiRes.AccessToken = util.UUIDStr()
	apiRes.ExpireIn = 3600

	return h.RespJson(c, apiRes, err)
}
