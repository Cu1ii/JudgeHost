package database

import (
	"JudgeHost/xoj/dao"
	"fmt"
	"testing"
)

func TestSelectJudgeStatus(t *testing.T) {
	all := []*dao.JudgeStatus{}
	mySQLDB := GetMySQLDB()
	if find := mySQLDB.Raw("SELECT * FROM judgestatus_judgestatus").Scan(&all); find.Error != nil {
		fmt.Println("find all judge_status fail")
	}
	for _, stauts := range all {
		fmt.Println(stauts)
	}
}

func TestSelectProblem(t *testing.T) {
	all := []*dao.Problem{}
	mySQLDB := GetMySQLDB()
	if find := mySQLDB.Raw("SELECT * FROM problem_problem").Scan(&all); find.Error != nil {
		fmt.Println("find all problem fail")
	}
	for _, stauts := range all {
		fmt.Println(stauts)
	}
}

func TestSelectProblemData(t *testing.T) {
	all := []*dao.ProblemData{}
	mySQLDB := GetMySQLDB()
	if find := mySQLDB.Raw("SELECT * FROM problem_problemdata").Scan(&all); find.Error != nil {
		fmt.Println("find all problem fail")
	}
	for _, stauts := range all {
		fmt.Println(stauts)
	}
}
