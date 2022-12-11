package database

import (
	"JudgeHost/xoj/constent"
	"JudgeHost/xoj/dao"
	"fmt"
	"github.com/sirupsen/logrus"
)

//---------------------------------JudgeStatus-----------------------------------------//

func GetJudgeStatus() []*dao.JudgeStatus {
	mySQLDB := GetMySQLDB()
	judgeArry := []*dao.JudgeStatus{}
	if res := mySQLDB.Raw("SELECT * FROM judgestatus_judgestatus where result = ?", constent.PENDING).Scan(&judgeArry); res.Error != nil {
		logrus.Error("select judge status error ", res.Error)
		return nil
	}
	return judgeArry
}

func GetJudgeStatusById(id int64) *dao.JudgeStatus {
	mySQLDB := GetMySQLDB()
	status := dao.JudgeStatus{}
	if res := mySQLDB.Raw("SELECT * FROM judgestatus_judgestatus where id = ?", id).Scan(&status); res.Error != nil {
		logrus.Error("select judge status error ", res.Error)
		return nil
	}
	return &status
}

func UpdateJudgeStatusResult(id int, result int) bool {
	mySQLDB := GetMySQLDB()
	if res := mySQLDB.Exec("UPDATE judge_backend.judgestatus_judgestatus SET result = ? WHERE id = ?", result, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

func UpdateJudgeStatusMessage(id int, msg string) bool {
	mySQLDB := GetMySQLDB()
	if res := mySQLDB.Exec("UPDATE judge_backend.judgestatus_judgestatus SET message = ? WHERE id = ?", msg, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

func UpdateJudgeStatus(id int, memory, mytime int, result int, testcase string) bool {
	mySQLDB := GetMySQLDB()
	if res := mySQLDB.Exec("UPDATE judgestatus_judgestatus SET memory = ?, time= ?, result = ?, testcase=?  WHERE id = ?", memory, mytime, result, testcase, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

//---------------------------------Problem-----------------------------------------//

func GetProblemById(pk string) *dao.Problem {
	mySQLDB := GetMySQLDB()
	problem := dao.Problem{}
	if res := mySQLDB.Raw("SELECT * FROM problem_problem WHERE problem = ?", pk).Scan(&problem); res.Error != nil {
		logrus.Error("select problem error ", res.Error)
		return nil
	}
	return &problem
}

func GetIsHaveDoneProblem(username, problem string) bool {
	mySQLDB := GetMySQLDB()
	selectProblem := fmt.Sprintf("SELECT * FROM judgestatus_judgestatus WHERE user = '%s'  AND problem = '%s' AND result = 0", username, problem)
	problems := []dao.Problem{}
	if res := mySQLDB.Raw(selectProblem).Scan(&problems); res.Error != nil {
		logrus.Error("select problem error ", res.Error)
		return false
	}
	if len(problems) > 0 {
		return true
	}
	return false
}

func AddProSubmitNum(problem string) bool {
	mySQLDB := GetMySQLDB()
	addProSubmitNum := fmt.Sprintf("UPDATE problem_problemdata SET submission = submission+1 WHERE problem = '%s'", problem)
	if res := mySQLDB.Exec(addProSubmitNum); res.Error != nil {
		logrus.Error("add problem data ( submission = submission + 1 ) error ", res.Error)
		return false
	}
	return true
}

func GetProblemTimeMemory(pk string) (int, int) {
	problem := GetProblemById(pk)
	return problem.Time, problem.Memory
}

func GetProblemScore(pk string) int {
	problemData := GetProblemDataById(pk)
	return problemData.Score
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
	if res := mySQLDB.Exec("UPDATE problem_problemdata SET submission = submission + 1, "+
		result+" = "+result+" + 1"+" WHERE problem = ?", pk); res.Error != nil {
		logrus.Error("update problem data submission error ", res.Error)
		return false
	}
	return true
}

func UpdateProblemAuth(pk string, auth int) bool {
	mySQLDB := GetMySQLDB()
	if res := mySQLDB.Exec("UPDATE  problem_problem SET auth = ? WHERE problem = ?", auth, pk); res.Error != nil {
		logrus.Error("update problem data auth error ", res.Error)
		return false
	}
	return true
}

func UpdateProblemDataAuth(pk string, auth int) bool {
	mySQLDB := GetMySQLDB()
	if res := mySQLDB.Exec("UPDATE  problem_problemdata SET auth = ? WHERE problem = ?", auth, pk); res.Error != nil {
		logrus.Error("update problem data auth error ", res.Error)
		return false
	}
	return true
}

//---------------------------------CaseStatus-----------------------------------------//

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

//---------------------------------Contest-----------------------------------------//

func SetBoard(id int64, statue int) bool {
	mySQLDB := GetMySQLDB()
	updateBoardType := fmt.Sprintf("UPDATE contest_contestboard SET type = %d WHERE submitid = %d", statue, id)
	if res := mySQLDB.Exec(updateBoardType); res.Error != nil {
		logrus.Error("update board type error ", res.Error)
		return false
	}
	return true
}

func GetNotExpiredContest() []*dao.ConstInfo {
	mySQLDB := GetMySQLDB()
	var data []*dao.ConstInfo
	res := mySQLDB.Raw("SELECT * from contest_contestinfo where type <> 'Personal' and TO_SECONDS(NOW()) - TO_SECONDS(begintime) <= lasttime").Scan(&data)
	if res.Error != nil {
		logrus.Error("select not expired contest error ", res.Error)
		return nil
	}
	return data
}

func GetRunningContest() []*dao.ConstInfo {
	mySQLDB := GetMySQLDB()
	var data []*dao.ConstInfo
	res := mySQLDB.Raw(
		"SELECT * from contest_contestinfo where type <> 'Personal' and " +
			"TO_SECONDS(NOW()) - TO_SECONDS(begintime) <= lasttime and TO_SECONDS(NOW()) - TO_SECONDS(begintime) >=-1").Scan(&data)
	if res.Error != nil {
		logrus.Error("select not expired contest error ", res.Error)
		return nil
	}
	return data
}

func GetContestProblem(contestId int) []*dao.Problem {
	mySQLDB := GetMySQLDB()
	problems := []*dao.Problem{}
	if res := mySQLDB.Raw("SELECT * from contest_contestproblem where contestid= ?", contestId).Scan(&problems); res.Error != nil {
		logrus.Error("select contest problem error ", res.Error)
		return nil
	}
	return problems
}

func UpdateContestBoardTypeBySubmitId(typ, id int) bool {
	mySQLDB := GetMySQLDB()
	if res := mySQLDB.Exec("UPDATE contest_contestboard SET type = ?  WHERE submitid = ?", typ, id); res.Error != nil {
		logrus.Error("update contest board type error ", res.Error)
		return false
	}
	return true
}

//---------------------------------User-----------------------------------------//

func UpdateUserResult(username, result string) bool {
	mySQLDB := GetMySQLDB()
	updateSQL := fmt.Sprintf("UPDATE problem_problemdata SET %s = %s + 1 WHRER username = %s", result, result, username)
	if res := mySQLDB.Exec(updateSQL); res.Error != nil {
		logrus.Error("update user result error ", res.Error)
		return false
	}
	return true
}

func UpdateUserScore(username string, score int) bool {
	mySQLDB := GetMySQLDB()
	updateSQL := fmt.Sprintf("UPDATE user_userdata SET score = score+%d WHERE username = '%s'", score, username)
	if res := mySQLDB.Exec(updateSQL); res.Error != nil {
		logrus.Error("update user score error ", res.Error)
		return false
	}
	return true
}

func UpdateUserAcPro(problem, username string) bool {
	mySQLDB := GetMySQLDB()
	updateSQL := fmt.Sprintf("UPDATE user_userdata SET acpro = concat(acpro,'|%s') WHERE username = '%s'", problem, username)
	if res := mySQLDB.Exec(updateSQL); res.Error != nil {
		logrus.Error("update user score error ", res.Error)
		return false
	}
	return true
}
