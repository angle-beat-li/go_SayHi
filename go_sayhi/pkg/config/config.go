package config

import (
	"strings"

	"github.com/mlogclub/simple/sqls"
)

var Instance *Config

type Config struct {
	Env            string   // 环境
	BaseUrl        string   // base url
	Port           string   // 端口
	IpDatePath     string   // IP数据文件
	AllowedOrigins []string // 跨域白名单

	// 日志配置
	Logger struct {
		Filename   string // 日志文件的位置
		MaxSize    int    // 文件最大尺寸（以MB为单位）
		MaxAge     int    // 保留旧文件的最大天数
		MaxBackups int    // 保留的最大旧文件数量
	}

	// 数据库配置
	DB sqls.DbConfig

	//smtp
	Smtp struct {
		Host     string
		Port     string
		Username string
		Password string
		SSL      bool
	}

	Search struct {
		IndexPath string
	}
}

func (c *Config) IsProd() bool {
	e := strings.ToLower(c.Env)
	return e == "prod" || e == "production"
}
