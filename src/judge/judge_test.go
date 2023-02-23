package judge

import (
	"fmt"
	"testing"
)

func TestJudgeCPP(t *testing.T) {
	res, _ := judgeCPP(
		1000,
		64,
		1,
		"/home/cu1/XOJ/resolutions/1/0.in",
		"/home/cu1/XOJ/submission/1/0temp.out",
		"/home/cu1/XOJ/submission/1/1error.out",
		"/home/cu1/XOJ/submission/1/1log.out",
	)
	fmt.Println(res)
}

func TestJudgeC(t *testing.T) {
	res, _ := judgeC(
		1000,
		64,
		2,
		"/home/cu1/XOJ/resolutions/1/0.in",
		"/home/cu1/XOJ/submission/2/0temp.out",
		"/home/cu1/XOJ/submission/2/2error.out",
		"/home/cu1/XOJ/submission/2/2log.out",
	)
	fmt.Println(res)
}

func TestJudgeGo(t *testing.T) {
	res, _ := judgeGo(
		1000,
		64,
		5,
		"/home/cu1/XOJ/resolutions/1/0.in",
		"/home/cu1/XOJ/submission/5/0temp.out",
		"/home/cu1/XOJ/submission/5/5error.out",
		"/home/cu1/XOJ/submission/5/5log.out",
	)
	fmt.Println(res)
}

func TestJudgeJava(t *testing.T) {
	res, _ := judgeJava(
		1000,
		64,
		6,
		"/home/cu1/XOJ/resolutions/1/0.in",
		"/home/cu1/XOJ/submission/6/0temp.out",
		"/home/cu1/XOJ/submission/6/6error.out",
		"/home/cu1/XOJ/submission/6/6log.out",
	)
	fmt.Println(res)
}

func TestJudgePython2(t *testing.T) {
	res, _ := judgePyhton2(
		1000,
		64,
		3,
		"/home/cu1/XOJ/resolutions/1/0.in",
		"/home/cu1/XOJ/submission/3/0temp.out",
		"/home/cu1/XOJ/submission/3/3error.out",
		"/home/cu1/XOJ/submission/3/3log.out",
	)
	fmt.Println(res)
}

func TestJudgePython3(t *testing.T) {
	res, _ := judgePyhton3(
		1000,
		64,
		4,
		"/home/cu1/XOJ/resolutions/1/0.in",
		"/home/cu1/XOJ/submission/4/0temp.out",
		"/home/cu1/XOJ/submission/4/4error.out",
		"/home/cu1/XOJ/submission/4/4log.out",
	)
	fmt.Println(res)
}
