package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"magic_box/server/entity"
	"magic_box/server/handler"
	"magic_box/server/service"
)

type (
	Server struct {
		cfg entity.Config
		e   *echo.Echo
		svc *service.Service
	}
)

func NewServer(cfg entity.Config, svc *service.Service, logWriter io.Writer) (*Server, error) {
	s := &Server{
		cfg: cfg,
		svc: svc,
	}

	e := echo.New()
	e.Binder = &Binder{}
	e.HTTPErrorHandler = handler.ErrorHandler
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper:          middleware.DefaultLoggerConfig.Skipper,
		Format:           middleware.DefaultLoggerConfig.Format,
		CustomTimeFormat: middleware.DefaultLoggerConfig.CustomTimeFormat,
		CustomTagFunc:    middleware.DefaultLoggerConfig.CustomTagFunc,
		Output:           logWriter,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))
	e.Use(middleware.Recover())
	e.Pre(Pre)

	s.e = e

	return s, nil
}

func (s *Server) Run() error {
	s.initRoute()

	return s.e.Start(fmt.Sprintf("%s:%d", s.cfg.Http.Addr, s.cfg.Http.Port))
}

func (s *Server) initRoute() {
	if s == nil || s.e == nil {
		return
	}
	initRoute(s.e, s.svc)
}
