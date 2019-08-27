package config

import (
	"github.com/haozzzzzzzz/go-rapid-development/utils/yaml"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

type ServiceConfigFormat struct {
	Host string `yaml:"host" validate:"required"`
	Port string `yaml:"port" validate:"required"`
}

var ServiceConfig *ServiceConfigFormat

func CheckServiceConfig() {
	if ServiceConfig != nil {
		return
	}

	ServiceConfig = &ServiceConfigFormat{}

	err := yaml.ReadYamlFromFile("./config/service.yaml", ServiceConfig)
	if nil != err {
		logrus.Fatalf("read service config failed. %s", err)
		return
	}

	err = validator.New().Struct(ServiceConfig)
	if err != nil {
		logrus.Fatalf("validate service config failed. %s", err)
	}

}