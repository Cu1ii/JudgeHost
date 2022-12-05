package database

import (
	"JudgeHost/xoj/dao"
	"fmt"
	"testing"
)

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
