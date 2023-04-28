package repository

import (
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"magic_box/server/repository/ent"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type (
	DatabaseConfig struct {
		Schema               string
		Addr                 string
		Port                 uint16
		SSL                  bool
		Username             string
		Password             string
		Database             string
		Charset              string
		Debug                bool
		MaxConnLimit         int
		IdleConnLimit        int
		MaxIdleTimeoutSecond uint
		TimeoutSecond        uint
	}
)

func (cfg DatabaseConfig) ToDataSource() string {
	var (
		ds     string
		values = url.Values{}
	)

	switch cfg.Schema {
	case "postgres":
		ds = fmt.Sprintf("%s://%s:%s@%s:%d/%s",
			cfg.Schema, cfg.Username, cfg.Password,
			cfg.Addr, cfg.Port, cfg.Database)
		if len(cfg.Charset) > 0 {
			values.Set("client_encoding", cfg.Charset)
		}
		if cfg.SSL {
			// 这里只支持ca
			values.Set("sslmode", "verify-ca")
		} else {
			values.Set("sslmode", "disable")
		}
		if cfg.TimeoutSecond > 0 {
			values.Set("connect_timeout", fmt.Sprintf("%d", cfg.TimeoutSecond))
		}
	case "mysql":
		ds = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			cfg.Username, cfg.Password,
			cfg.Addr, cfg.Port, cfg.Database)
		values.Set("parseTime", "true")
		if len(cfg.Charset) > 0 {
			values.Set("charset", cfg.Charset)
		}
		if cfg.SSL {
			values.Set("tls", "true")
		} else {
			values.Set("tls", "false")
		}
		if cfg.TimeoutSecond > 0 {
			values.Set("timeout", fmt.Sprintf("%ds", cfg.TimeoutSecond))
		}
	default:
		panic("not support database schema:" + cfg.Schema)
	}
	if len(values.Encode()) > 0 {
		ds = ds + "?" + values.Encode()
	}

	return ds
}

func New(cfg DatabaseConfig) (*ent.Client, error) {
	db, err := sql.Open(cfg.Schema, cfg.ToDataSource())
	if err != nil {
		return nil, err
	}
	// 设置参数
	db.SetMaxIdleConns(cfg.MaxConnLimit)
	db.SetMaxOpenConns(cfg.IdleConnLimit)
	db.SetConnMaxLifetime(time.Second * time.Duration(cfg.TimeoutSecond))
	db.SetConnMaxIdleTime(time.Second * time.Duration(cfg.MaxIdleTimeoutSecond))

	// 获取客户端
	ent.Log(func(a ...any) {
		log.Infoln(a...)
	})
	cli := ent.NewClient(ent.Driver(entsql.OpenDB(cfg.Schema, db)))
	if cfg.Debug {
		return cli.Debug(), nil
	}
	return cli, nil
}
