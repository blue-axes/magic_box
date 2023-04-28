package user

import (
	"magic_box/server/handler"
	"magic_box/server/service"
)

type (
	User struct {
		*handler.Handler
		svc *service.Service
	}
)

func NewUser(svc *service.Service) *User {
	h := &User{
		Handler: handler.New(svc),
		svc:     svc,
	}
	return h
}
