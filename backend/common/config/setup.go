package config

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"

	"github.com/sirupsen/logrus"
)

func setupLogPanic() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(ServiceConfig.LogLevel)

	if ServiceConfig.LogOutput != nil {
		configOutput := ServiceConfig.LogOutput
		fileName := fmt.Sprintf("%s/%s/info.log", configOutput.Filedir, ServiceConfig.ServiceName)
		logger := &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    configOutput.MaxSize,
			MaxAge:     configOutput.MaxAge,
			MaxBackups: configOutput.MaxBackups,
			LocalTime:  false,
			Compress:   configOutput.Compress,
		}

		logrus.SetOutput(logger)
		logrus.SetOutput(logger)

	} else { // std
		logger := logrus.New()
		log.SetOutput(logger.Writer())

	}

}
