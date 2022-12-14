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
		for _, status := range pendingStatus {
			database.UpdateJudgeStatusResult(int(status.Id), constent.WAITING)
			database.UpdateJudgeStatusJudger(int(status.Id), "XOJ")
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
				if _, ok := curPro[pro.ProblemId]; !ok {
					logrus.Info("58 pro.Problem", pro.ProblemId)
					curPro[pro.ProblemId] = true
					database.UpdateProblemDataAuth(pro.ProblemId, 2)
					database.UpdateProblemAuth(pro.ProblemId, 2)
				}
			}
		}
		runningContests := database.GetRunningContest()
		for _, runningContest := range runningContests {
			contestProblems := database.GetContestProblem(runningContest.Id)
			for _, pro := range contestProblems {
				if _, ok := curRunPro[pro.ProblemId]; !ok {
					curRunPro[pro.ProblemId] = true
					database.UpdateProblemDataAuth(pro.ProblemId, 3)
					database.UpdateProblemAuth(pro.ProblemId, 3)
				}
			}
		}
		for contestId, _ := range curContest {
			if _, ok := allContest[contestId]; !ok {
				contestProblems := database.GetContestProblem(contestId)
				for _, pro := range contestProblems {
					if _, ok := curPro[pro.ProblemId]; ok {
						delete(curPro, pro.ProblemId)
					}
					if _, ok := curRunPro[pro.ProblemId]; ok {
						delete(curRunPro, pro.ProblemId)
					}
					database.UpdateProblemDataAuth(pro.ProblemId, 1)
					database.UpdateProblemAuth(pro.ProblemId, 1)
				}
			}
		}
		curContest = allContest
	}
}
