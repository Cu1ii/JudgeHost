package configuration

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"sync/atomic"
)

type Task interface {
	Do()
}

var JudgeIndexId int64 = 1
var JudgeNamePrefix string

func SetJudgeNamePrefix(name string) {
	JudgeNamePrefix = "From JudgeThreadFactory's " + name + "-Worker-"
}

type TaskWrop struct {
	Task Task
	Name string
	Wg   *sync.WaitGroup
}

func NewTaskWrop(task Task, wg *sync.WaitGroup) *TaskWrop {
	name := JudgeNamePrefix + strconv.FormatInt(atomic.AddInt64(&JudgeIndexId, 1), 10)
	return &TaskWrop{Task: task, Name: name, Wg: wg}
}

func TaskFunc(data interface{}) {
	taskWrop := data.(*TaskWrop)
	taskWrop.Task.Do()
	logrus.Info("taskWrop.Name: ", taskWrop.Name)
	taskWrop.Wg.Done()
}
