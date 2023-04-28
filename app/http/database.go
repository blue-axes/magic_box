package main

import (
	"magic_box/server/entity"
	"magic_box/server/repository"
	"magic_box/server/repository/ent"
)

func GetRepo(cfg entity.Config, retryTimes int) (*ent.Client, error) {
	repo, err := repository.New(repository.DatabaseConfig{
		Schema:               cfg.Database.Schema,
		Addr:                 cfg.Database.Addr,
		Port:                 cfg.Database.Port,
		SSL:                  cfg.Database.SSL,
		Username:             cfg.Database.Username,
		Password:             cfg.Database.Password,
		Database:             cfg.Database.Database,
		Charset:              cfg.Database.Charset,
		Debug:                cfg.Database.Debug,
		MaxConnLimit:         cfg.Database.MaxConnLimit,
		IdleConnLimit:        cfg.Database.IdleConnLimit,
		MaxIdleTimeoutSecond: cfg.Database.MaxIdleTimeoutSecond,
		TimeoutSecond:        cfg.Database.TimeoutSecond,
	})
	if err != nil {
		if retryTimes < 0 { // 一直重试
			return GetRepo(cfg, retryTimes)
		} else if retryTimes > 0 { // 重试次数大于0
			return GetRepo(cfg, retryTimes-1)
		} else { // 重试次数=0
			return nil, err
		}
	}
	return repo, nil
}
