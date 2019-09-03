package db

import (
	"backend/common/config"
	"github.com/haozzzzzzzz/go-rapid-development/db"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

var HubSqlxDB *sqlx.DB
var HubDBConfig *db.ClientConfigFormat

func init() {
	var err error
	HubDBConfig = &config.DBConfig
	HubSqlxDB, err = db.NewDB(HubDBConfig)
	if nil != err {
		logrus.Fatalf("new db %s connection failed. %s", HubDBConfig.Source, err)
		return
	}
}
