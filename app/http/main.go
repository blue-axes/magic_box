package main

import (
	"context"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"magic_box/pkg/log_hook"
	"magic_box/server/entity"
	"magic_box/server/repository/ent/migrate"
	"magic_box/server/service"
	"os"
	"path"
)

func main() {
	// 解析命令行参数
	args := parseArgs()
	// 加载配置文件
	cfg, err := entity.LoadConfig(args.Config, true)
	if err != nil {
		fmt.Println("load config error:" + err.Error())
		os.Exit(-1)
	}
	// 初始化日志组件
	logWriter, err := initLogComponent(cfg)
	if err != nil {
		fmt.Println("init log component error:" + err.Error())
		os.Exit(-1)
	}

	// 初始化数据库连接以及初始化数据库
	repo, err := GetRepo(cfg, -1) // 无限重试
	if err != nil {
		fmt.Println("connect database error:" + err.Error())
		os.Exit(-1)
	}
	err = repo.Schema.Create(context.Background(), migrate.WithForeignKeys(true))
	if err != nil {
		fmt.Println("initialize database tables error:" + err.Error())
		os.Exit(-1)
	}

	// 获取svc
	svc, err := service.New(service.Config{
		Repo: repo,
	})
	if err != nil {
		fmt.Println("initialize service error:" + err.Error())
		os.Exit(-1)
	}

	svr, err := NewServer(cfg, svc, logWriter)
	if err != nil {
		fmt.Println("initialize server error:" + err.Error())
		os.Exit(-1)
	}
	fmt.Println(svr.Run())
}

func initLogComponent(cfg entity.Config) (io.Writer, error) {
	log.SetFormatter(&log.JSONFormatter{})
	log.AddHook(log_hook.New())
	lvl, err := log.ParseLevel(cfg.SystemConf.LogLevel)
	if err != nil {
		lvl = log.InfoLevel
	}
	options := make([]rotatelogs.Option, 0)
	options = append(options, rotatelogs.WithLinkName(cfg.SystemConf.LogFilename))
	if cfg.SystemConf.LogFileMaxSizeByte > 0 {
		options = append(options, rotatelogs.WithRotationSize(int64(cfg.SystemConf.LogFileMaxSizeByte)))
	}
	if cfg.SystemConf.LogFileCount > 0 {
		options = append(options, rotatelogs.WithRotationCount(uint(cfg.SystemConf.LogFileCount)))
	}

	log.SetLevel(lvl)
	_ = os.MkdirAll(path.Dir(cfg.SystemConf.LogFilename), 0666)
	w, err := rotatelogs.New(cfg.SystemConf.LogFilename+"-%Y%m%d%H%M", options...)
	if err != nil {
		return nil, err
	}
	log.SetOutput(w)
	return w, nil
}
