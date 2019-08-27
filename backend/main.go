package main

import (
	"backend/api"
	"backend/config"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/haozzzzzzzz/go-rapid-development/web/ginbuilder"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	var err error
	serviceConfig := config.ServiceConfig
	address := fmt.Sprintf("%s:%s", serviceConfig.Host, serviceConfig.Port)

	engine := ginbuilder.DefaultEngine()
	err = api.BindRouters(engine)
	if nil != err {
		logrus.Errorf("bind routers failed. error: %s.", err)
		return
	}

	endless.DefaultReadTimeOut = 10 * time.Second
	endless.DefaultWriteTimeOut = 10 * time.Second
	err = endless.ListenAndServe(address, engine)
	if nil != err {
		logrus.Errorf("start listening and serving http on %s failed. error: %s.", address, err)
		return
	}
}
