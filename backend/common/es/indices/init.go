package indices

import (
	"backend/common/config"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/haozzzzzzzz/go-rapid-development/es"
	"github.com/sirupsen/logrus"
)

var EsClient *elasticsearch.Client

func init() {
	var err error

	EsClient, err = es.NewClient(
		[]string{config.ElasticSearchConfig.Endpoint},
		es.ShortTimeoutTransport,
	)
	if nil != err {
		logrus.Panicf("new es client failed. %s", err)
		return
	}

}
