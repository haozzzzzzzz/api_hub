package sched

import (
	"backend/common/business"
	"context"
	"github.com/sirupsen/logrus"
)

func EsIndexAllAhDoc(ctx context.Context) (err error) {
	bsDoc := business.NewBsDoc(ctx)
	err = bsDoc.EsIndexAllAhDoc()
	if nil != err {
		logrus.Errorf("index all doc failed. error: %s.", err)
		return
	}
	return
}
