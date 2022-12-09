package xoj

import (
	"JudgeHost/xoj/database"
	"JudgeHost/xoj/pool"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

func JudgeSplitMode() {

	for true {
		time.Sleep(time.Second * 2)
		pendingStatus := database.GetJudgeStatus()
		logrus.Info("pending status = ", len(pendingStatus))
		wg := sync.WaitGroup{}
		for _, status := range pendingStatus {
			wg.Add(1)
			if err := pool.GetJudgePool().Submit(
				func() {
					judge(
						int(status.Id),
						status.Code,
						status.Language,
						status.Problem,
						int(status.Contest),
						status.User,
						status.Oj,
						"XOJ",
						status.SubmitTime,
						status.ContestProblem,
						false,
					)
				}); err != nil {
				logrus.Error("run judge ", status.Problem, " the user is ", status.User, " error: ", err.Error())
			}
		}
	}
}
