package database

import (
	"JudgeHost/xoj/dao"
	"github.com/sirupsen/logrus"
)

const pending = -1

func GetJudgeStatus() []*dao.JudgeStatus {
	mySQLDB := GetMySQLDB()
	judgeArry := []*dao.JudgeStatus{}
	if res := mySQLDB.Raw("SELECT * FROM judgestatus_judgestatus where result = ?", pending).Scan(&judgeArry); res.Error != nil {
		logrus.Error("select judge status error ", res.Error)
		return nil
	}
	return judgeArry
}

func GetProblemById(pk string) *dao.Problem {
	mySQLDB := GetMySQLDB()
	problem := dao.Problem{}
	if res := mySQLDB.Raw("SELECT * FROM problem_problem WHERE problem = ?", pk).Scan(&problem); res.Error != nil {
		logrus.Error("select problem error ", res.Error)
		return nil
	}
	return &problem
}

func GetProblemDataById(pk string) *dao.ProblemData {
	mySQLDB := GetMySQLDB()
	problemData := dao.ProblemData{}
	if res := mySQLDB.Raw("SELECT * FROM problem_problemdata WHERE problem = ?", pk).Scan(&problemData); res.Error != nil {
		logrus.Error("select problem data error ", res.Error)
		return nil
	}
	return &problemData
}

func UpdateProblemData(pk string, result string) bool {
	mySQLDB := GetMySQLDB()
	if res := mySQLDB.Raw("UPDATE problem_problemdata SET submission = submission + 1 "+
		result+" = "+result+" + 1"+"WHERE problem = ? ESCAPE '#'", pk); res.Error != nil {
		logrus.Error("update problem data error ", res.Error)
		return false
	}
	return true
}

func AddCaseStatus(status *dao.CaseStatus) bool {
	mySQLDB := GetMySQLDB()
	if create := mySQLDB.Exec("INSERT INTO judgestatus_casestatus "+
		"(statusid, username, problem, result, time, memory, testcase, casedata, outputdata, useroutput)"+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		status.StatusId, status.Username, status.Problem, status.Result,
		status.Time, status.Memory, status.TestCase, status.CaseData, status.OutputData, status.UserOutput); create.Error != nil {
		logrus.Error("insert case_status error ", create.Error)
		return false
	}
	return true
}