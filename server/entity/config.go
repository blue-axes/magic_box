package entity

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		Http       HttpConfig     `json:"http"`
		Database   DatabaseConfig `json:"database"`
		InitData   InitData       `json:"init_data"`
		SystemConf SystemConf     `json:"system_conf"`
	}
	HttpConfig struct {
		Addr string `json:"addr"`
		Port uint16 `json:"port"`
	}
	DatabaseConfig struct {
		Schema   string `json:"schema"`
		Addr     string `json:"addr"`
		Port     uint16 `json:"port"`
		SSL      bool   `json:"ssl"`
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
		Charset  string `json:"charset"`
		Debug    bool   `json:"debug"`

		MaxConnLimit         int  `json:"max_conn_limit"`
		IdleConnLimit        int  `json:"idle_conn_limit"`
		MaxIdleTimeoutSecond uint `json:"max_idle_timeout_second"`
		TimeoutSecond        uint `json:"timeout_second"`
	}

	// InitData 初始化的数据
	InitData struct {
		Authentications []AuthenticationInfo `json:"authentications"`
	}
	AuthenticationInfo struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	SystemConf struct {
		LogLevel           string `json:"log_level"`
		LogFilename        string `json:"log_filename"`
		LogFileMaxSizeByte uint64 `json:"log_file_max_size_byte"`
		LogFileCount       int    `json:"log_file_count"`
	}
)

var (
	// 配置缓存
	__config_cache__ *Config
)

func LoadConfig(cfgPath string, reloadForce bool) (cfg Config, err error) {
	if !reloadForce && __config_cache__ != nil {
		return *__config_cache__, nil
	}
	f, err := os.OpenFile(cfgPath, os.O_RDONLY, 0666)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	dec.UseNumber()
	err = dec.Decode(&cfg)
	return cfg, err
}
