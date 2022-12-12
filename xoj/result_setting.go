package xoj

import (
	"JudgeHost/xoj/dao"
	"JudgeHost/xoj/database"
	"github.com/sirupsen/logrus"
)

func compileError(id int, problem, msg string) {
	logrus.Info("Compile error! ", id)
	database.UpdateJudgeStatusResult(id, -4)
	database.UpdateJudgeStatusMessage(id, msg)
	database.UpdateProblemData(problem, "ce")
}

func doneProblem(id int, problem, message string, memory, mytime int, username string, contest, result int, testcase string) {
	// fmt.Println(id, " ", problem, " ", message, " ", mytime, " ", username, " ", contest, " ", result, " ", testcase)
	database.UpdateJudgeStatus(id, memory, mytime, result, testcase)
	if message != "" {
		database.UpdateJudgeStatusMessage(id, message)
	}
	if result == 2 || result == 1 {
		database.UpdateProblemData(problem, "tle")
	}
	if result == 3 {
		database.UpdateProblemData(problem, "mle")
	}
	if result == 4 {
		database.UpdateProblemData(problem, "rte")
	}
	if result == 5 {
		database.UpdateProblemData(problem, "se")
	}
	if result == -5 {
		database.UpdateProblemData(problem, "pe")
	}
	if result == -3 {
		database.UpdateProblemData(problem, "wa")
	}
	if contest != 0 {
		database.UpdateContestBoardTypeBySubmitId(0, id)
	}
	database.UpdateUserResult(username, "submit")
}

func acProblem(id int, problem, message string, memory, time int, username string, proScore int, isAc bool, contest int) {
	//fmt.Println(time, " ", memory)
	database.UpdateJudgeStatus(id, memory, time, 0, "")
	if message != "" {
		database.UpdateJudgeStatusMessage(id, message)
	}
	database.UpdateProblemData(problem, "ac")
	if !isAc {
		database.UpdateUserResult(username, "ac")
		database.UpdateUserScore(username, proScore)
		database.UpdateUserAcPro(problem, username)
	}
	if contest != 0 {
		database.UpdateContestBoardTypeBySubmitId(1, id)
	}
	database.UpdateUserResult(username, "submit")
}

func doneCase(statusId int, username, problem, result string,
	time, memory int, testcase, caseData, outputData, userOutput string) {
	database.AddCaseStatus(&dao.CaseStatus{
		StatusId:   statusId,
		Username:   username,
		Problem:    problem,
		Result:     result,
		Time:       time,
		Memory:     memory,
		TestCase:   testcase,
		CaseData:   caseData,
		OutputData: outputData,
		UserOutput: userOutput,
	})
}
