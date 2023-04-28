package handler

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"magic_box/pkg/errors"
	"magic_box/server/entity"
	"magic_box/server/service"
	"net/http"
)

type (
	Handler struct {
		svc *service.Service
	}
	respStruct struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		TraceID string      `json:"trace_id"`
		Data    interface{} `json:"data"`
	}
)

func New(svc *service.Service) *Handler {
	h := &Handler{
		svc: svc,
	}
	return h
}

func (h Handler) Ctx(c echo.Context) *entity.Context {
	traceID, _ := c.Get(entity.CtxTraceID).(string)
	ctx := entity.NewContext(traceID)
	return ctx
}

func (h Handler) Log(ctx *entity.Context) *log.Entry {
	if ctx == nil {
		return log.WithField("trace_id", "")
	}
	return log.WithField("trace_id", ctx.TraceID)
}

func (h Handler) RespJson(c echo.Context, data interface{}, err error) error {
	if err != nil {
		ErrorHandler(err, c)
		return nil
	}
	traceID, _ := c.Get("CtxTraceID").(string)

	resp := respStruct{
		Code:    entity.CodeOk,
		Message: "",
		TraceID: traceID,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func ErrorHandler(err error, c echo.Context) {
	traceID, _ := c.Get("CtxTraceID").(string)
	resp := respStruct{
		Code:    entity.CodeUnexpected,
		Message: "",
		TraceID: traceID,
	}
	switch verr := err.(type) {
	case *errors.Error:
		resp.Code = verr.Code()
		resp.Message = verr.Message()
	default:
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
