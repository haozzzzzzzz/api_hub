package sched

import (
	"github.com/haozzzzzzzz/go-rapid-development/task"
	"github.com/sirupsen/logrus"
	"time"
)

var scheduler = task.New()

func RunSchedulerPanic() {
	var err error

	var boolLocker task.BoolLocker
	err = scheduler.AddTaskJob(&task.TaskJob{
		Spec:        "@hourly",
		TaskName:    "es_index_all_ah_doc",
		Handler:     EsIndexAllAhDoc,
		ExecTimeout: 1 * time.Hour,
		Locker:      &boolLocker,
	})
	if nil != err {
		logrus.Panicf("add task job failed. error: %s.", err)
		return
	}

	scheduler.Start()
}
