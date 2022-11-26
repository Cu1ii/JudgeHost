package configuration

import (
	"github.com/sirupsen/logrus"
	"sync/atomic"
)

type Task interface {
	Do()
	GetName() string
	Wait()
}

var JudgeIndexId atomic.Int64
var JudgeNamePrefix string

func SetJudgeNamePrefix(name string) {
	JudgeNamePrefix = "From JudgeThreadFactory's " + name + "-Worker-"
}

func TaskFunc(data interface{}) {
	task := data.(Task)
	logrus.Info("task.Name: ", task.GetName)
	task.Do()
}
