package vo

import "JudgeHost/src/models/dto"

// JudgeConditionVO 对某次判题最终结果的表现层对象
type JudgeConditionVO struct {
	JudgeResults []*dto.SingleJudgeResultDTO
	SubmissionId string
	JudgeEndTime int64
	ExtraInfo    []string
}

func NewJudgeConditionVO(judgeResults []*dto.SingleJudgeResultDTO,
	compileResult []string, submssionId string) *JudgeConditionVO {
	return &JudgeConditionVO{
		JudgeResults: judgeResults,
		SubmissionId: submssionId,
		ExtraInfo:    compileResult,
	}
}
