package config

import (
	"github.com/haozzzzzzzz/go-rapid-development/config"
	"github.com/haozzzzzzzz/go-rapid-development/db"
	"github.com/sirupsen/logrus"
)

type ServiceConfigFormat struct {
	ServiceName string `yaml:"service_name" validate:"required"`

	Host string `yaml:"host" validate:"required"`
	Port string `yaml:"port" validate:"required"`

	LogLevel  logrus.Level           `yaml:"log_level" validate:"required"`
	LogOutput *LogOutputConfigFormat `yaml:"log_output"`
}

type LogOutputConfigFormat struct {
	Filedir    string `json:"filedir" yaml:"filedir" validate:"required"`
	MaxSize    int    `json:"max_size" yaml:"max_size" validate:"required"`
	MaxBackups int    `json:"max_backups" yaml:"max_backups" validate:"required"`
	MaxAge     int    `json:"max_age" yaml:"max_age" validate:"required"`
	Compress   bool   `json:"compress" yaml:"compress"`
}

// 反向代理
type ReverseProxyConfigFormat struct {
	ProxyTypeMap map[string]string `yaml:"proxy_type_map"` // proxy_type => proxy_path_prefix
}

var ServiceConfig ServiceConfigFormat
var DBConfig db.ClientConfigFormat
var ReverseProxyConfig = ReverseProxyConfigFormat{
	ProxyTypeMap: make(map[string]string),
}

// panic if fail, for stopping process
func loadPanic() {
	config.LoadFileYamlPanic("./config/service.yaml", &ServiceConfig)
	config.LoadFileYamlPanic("./config/db.yaml", &DBConfig)
	config.LoadFileYamlNoError("./config/reverse_proxy.yaml", &ReverseProxyConfig)
}
