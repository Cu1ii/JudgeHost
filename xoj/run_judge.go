package xoj

import (
	"JudgeHost/xoj/database"
	"JudgeHost/xoj/pool"
	"github.com/sirupsen/logrus"
)

func JudgeSplitMode() {
	for true {
		pendingStatus := database.GetJudgeStatus()
		for _, status := range pendingStatus {
			if err := pool.GetJudgePool().Submit(
				func() {
					// TODO 编译

					// TODO 获取题目信息 TimeLimit, MemoryLimit, OutputLimit 等
					_ = database.GetProblemById(status.Problem)
					// TODO 判题

					// TODO 获取结果后更新表
					problemData := database.GetProblemDataById(status.Problem)
					result := GetResult()
					database.UpdateProblemData(problemData.Problem, result)
				}); err != nil {
				logrus.Error("run judge ", status.Problem, " the user is ", status.User, " error: ", err.Error())
			}
		}
	}
}

// GetResolutionPath TODO 获取 input, output 路径
func GetResolutionPath() {

}

// GetResult 获取最终结果映射 WRONG_ANSWER --> wa
func GetResult() string {
	return "nil"
}
