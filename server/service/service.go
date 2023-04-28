package service

import (
	log "github.com/sirupsen/logrus"
	"magic_box/server/entity"
	"magic_box/server/repository/ent"
)

type (
	Service struct {
		repo *ent.Client
	}
	Config struct {
		Repo *ent.Client
	}
)

func New(cfg Config) (*Service, error) {
	svc := &Service{
		repo: cfg.Repo,
	}

	// 服务层初始化
	err := svc.initialize()
	return svc, err
}

func (s Service) log(ctx entity.Context) *log.Entry {
	return log.WithField("trace_id", ctx.TraceID)
}
