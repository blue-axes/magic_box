package tools

import (
	"encoding/base64"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"magic_box/pkg/errors"
	"magic_box/server/entity"
	"magic_box/server/handler"
	"magic_box/server/handler/params"
	"magic_box/server/service"
	"net/url"
	"strings"
)

type (
	Tool struct {
		*handler.Handler
		svc *service.Service
	}
)

func NewTool(svc *service.Service) *Tool {
	h := &Tool{
		Handler: handler.New(svc),
		svc:     svc,
	}
	return h
}

func (h Tool) Base64(c echo.Context) error {
	var (
		apiReq = params.Base64EncodingReq{}
		apiRes = params.Base64EncodingResp{}
	)
	if err := c.Bind(&apiReq); err != nil {
		return err
	}
	var enc *base64.Encoding
	switch apiReq.Mode {
	case "std":
		fallthrough
	default:
		enc = base64.StdEncoding
	case "raw-std":
		enc = base64.RawStdEncoding
	case "raw-url":
		enc = base64.RawURLEncoding
	case "url":
		enc = base64.URLEncoding
	}

	switch apiReq.Action {
	case "encode":
		apiRes.Result = enc.EncodeToString([]byte(apiReq.Payload))
	case "decode":
		res, err := enc.DecodeString(apiReq.Payload)
		if err != nil {
			return errors.New(entity.CodeInvalidArgs, "invalid base64 data.")
		}
		apiRes.Result = string(res)
	}
	return h.RespJson(c, apiRes, nil)
}

func (h Tool) JsonPretty(c echo.Context) error {
	var (
		apiReq = params.JsonPrettyReq{}
		apiRes = params.JsonPrettyResp{}
	)
	if err := c.Bind(&apiReq); err != nil {
		return err
	}
	if apiReq.Indent != "" {
		apiReq.Indent = "\t"
	}

	payload := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(apiReq.Payload))
	dec.UseNumber()
	err := dec.Decode(&payload)
	if err != nil {
		return errors.New(entity.CodeInvalidArgs, "invalid json data."+err.Error())
	}
	res, _ := json.MarshalIndent(payload, "", apiReq.Indent)
	apiRes.Result = string(res)
	return h.RespJson(c, apiRes, nil)
}

func (h Tool) Url(c echo.Context) error {
	var (
		apiReq = params.UrlReq{}
		apiRes = params.UrlResp{}
		err    error
	)
	if err = c.Bind(&apiReq); err != nil {
		return err
	}
	switch apiReq.Action {
	case "encode":
		apiRes.Result = url.PathEscape(apiReq.Payload)
	case "decode":
		fallthrough
	default:
		apiRes.Result, err = url.PathUnescape(apiReq.Payload)
	}
	return h.RespJson(c, apiRes, err)
}
