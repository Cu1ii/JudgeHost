package database

import (
	"JudgeHost/xoj/dao"
	"fmt"
	"testing"
)

func TestSelectAllJudgeStatus(t *testing.T) {
	judgeStatus := GetJudgeAllStatus()
	for _, status := range judgeStatus {
		fmt.Println(status)
	}
}

func TestSelectJudgeArry(t *testing.T) {
	judgeStatus := GetJudgeStatus()
	for _, status := range judgeStatus {
		fmt.Println(status)
	}
}

func TestSelectProblemByPk(t *testing.T) {
	problem := GetProblemById("1")
	fmt.Println(problem)
}

func TestSelectProblemDataByPk(t *testing.T) {
	problemData := GetProblemDataById("1")
	fmt.Println(problemData)
}

func TestCreateCaseStatus(t *testing.T) {
	caseStatus := AddCaseStatus(&dao.CaseStatus{})
	fmt.Println("insert result = ", caseStatus)
}

func TestUpdateProblemData(t *testing.T) {
	res := UpdateProblemData("1", "ac")
	fmt.Println(res)
}

func TestUpdateJudgeStatusResult(t *testing.T) {
	res := UpdateJudgeStatusResult(1, 0)
	fmt.Println(res)
}

func TestUpdateJudgeStatusMessage(t *testing.T) {
	res := UpdateJudgeStatusMessage(1, "111")
	fmt.Println(res)
}

func TestSelectContestProblemByContestId(t *testing.T) {
	judgeStatus := GetContestProblem(4)
	fmt.Println(len(judgeStatus))
	for _, status := range judgeStatus {
		fmt.Println(status)
	}
}
