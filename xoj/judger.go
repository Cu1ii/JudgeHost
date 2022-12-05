package xoj

import (
	"JudgeHost/xoj/database"
)

func JudgeSplitMode() {
	for true {
		pendingStatus := database.GetJudgeStatus()
		for _, status := range pendingStatus {
			// TODO 未来使用线程池 ants
			go func() {
				// TODO 获取题目信息 TimeLimit, MemoryLimit, OutputLimit 等
				_ = database.GetProblemById(status.Problem)
				// TODO 判题

				// TODO 获取结果后更新表
				problemData := database.GetProblemDataById(status.Problem)
				result := GetResult()
				database.UpdateProblemData(problemData.Problem, result)
			}()
		}
	}
}

// GetResolutionPath TODO 获取 input, output 路径
func GetResolutionPath() {}

// GetResult 获取最终结果映射 WRONG_ANSWER --> wa
func GetResult() string {
	return "nil"
}
