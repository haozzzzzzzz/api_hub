package task

import (
	"backend/task/sched"
)

func InitTask() {
	sched.RunSchedulerPanic()
}
