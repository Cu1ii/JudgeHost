package xoj

import (
	"JudgeHost/xoj/database"
	"fmt"
	"testing"
)

func TestRunJudge(t *testing.T) {
	RunJudge()
}

func TestJudge(t *testing.T) {
	status := database.GetJudgeStatusById(1)
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
}

func TestSpj(t *testing.T) {
	i := specialjudge(
		"/home/cu1/XOJ/resolutions/2/checker",
		"/home/cu1/XOJ/resolutions/2/test1.in",
		"/home/cu1/XOJ/submission/test1.out",
		"/home/cu1/XOJ/resolutions/2/test1.out",
	)
	fmt.Println(i)
}
