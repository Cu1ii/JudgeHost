package xoj

import (
	"JudgeHost/xoj/constent"
	"JudgeHost/xoj/database"
	"JudgeHost/xoj/pool"
	"github.com/sirupsen/logrus"
	"time"
)

func RunJudge() {
	go changeAuth()
	for true {
		time.Sleep(time.Second * 2)
		pendingStatus := database.GetJudgeStatus()
		logrus.Info("pending status = ", len(pendingStatus))
		for _, status := range pendingStatus {
			database.UpdateJudgeStatusResult(int(status.Id), constent.WAITING)
		}
		for _, status := range pendingStatus {
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

// 比赛题目设置为auth=2,contest开始时，自动设置题目为auth=3，比赛结束自动设置auth=1
func changeAuth() {
	curContest := map[int]bool{}
	curPro := map[string]bool{}
	curRunPro := map[string]bool{}
	for true {
		time.Sleep(time.Second * 2)
		allContest := map[int]bool{}
		notExpiredContests := database.GetNotExpiredContest()
		for _, notExpiredContest := range notExpiredContests {
			allContest[notExpiredContest.Id] = true
			contestProblems := database.GetContestProblem(notExpiredContest.Id)
			for _, pro := range contestProblems {
				if _, ok := curPro[pro.Problem]; !ok {
					curPro[pro.Problem] = true
					database.UpdateProblemDataAuth(pro.Problem, 2)
					database.UpdateProblemAuth(pro.Problem, 2)
				}
			}
		}
		runningContests := database.GetRunningContest()
		for _, runningContest := range runningContests {
			contestProblems := database.GetContestProblem(runningContest.Id)
			for _, pro := range contestProblems {
				if _, ok := curRunPro[pro.Problem]; !ok {
					curRunPro[pro.Problem] = true
					database.UpdateProblemDataAuth(pro.Problem, 3)
					database.UpdateProblemAuth(pro.Problem, 3)
				}
			}
		}
		for contestId, _ := range curContest {
			if _, ok := allContest[contestId]; !ok {
				contestProblems := database.GetContestProblem(contestId)
				for _, pro := range contestProblems {
					if _, ok := curPro[pro.Problem]; ok {
						delete(curPro, pro.Problem)
					}
					if _, ok := curRunPro[pro.Problem]; ok {
						delete(curRunPro, pro.Problem)
					}
					database.UpdateProblemDataAuth(pro.Problem, 1)
					database.UpdateProblemAuth(pro.Problem, 1)
				}
			}
		}
		curContest = allContest
	}
}
