package xoj

import (
	"JudgeHost/xoj/database"
	"JudgeHost/xoj/pool"
	"github.com/sirupsen/logrus"
	"time"
)

func JudgeSplitMode() {
	for true {
		time.Sleep(time.Second * 2)
		pendingStatus := database.GetJudgeStatus()
		for _, status := range pendingStatus {
			if err := pool.GetJudgePool().Submit(
				func() {

				}); err != nil {
				logrus.Error("run judge ", status.Problem, " the user is ", status.User, " error: ", err.Error())
			}
		}
	}
}
