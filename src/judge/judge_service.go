package judge

import "context"

type JudgeServiceImpl struct{}

func (s *JudgeServiceImpl) Judge(context context.Context, rep *JudgeRequest) (*JudgeResponse, error) {
	return RunJudge(rep)
}
