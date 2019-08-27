package config

import (
	"github.com/sirupsen/logrus"
	"log"
)

func CheckLogConfig() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	newLogger := logrus.New()
	newLogger.Formatter = &logrus.JSONFormatter{}
	log.SetOutput(newLogger.Writer())
}