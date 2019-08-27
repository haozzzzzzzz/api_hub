package config

import (
	"github.com/haozzzzzzzz/go-rapid-development/db"
	"github.com/haozzzzzzzz/go-rapid-development/utils/yaml"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

var DBConfig *db.ClientConfigFormat

func CheckDBConfig() {
	if DBConfig != nil {
		return
	}

	DBConfig = &db.ClientConfigFormat{}
	err := yaml.ReadYamlFromFile("./config/db.yaml", DBConfig)
	if nil != err {
		logrus.Fatalf("read db config failed. %s", err)
		return
	}

	err = validator.New().Struct(DBConfig)
	if nil != err {
		logrus.Fatalf("validate db config failed. %s", err)
		return
	}
}