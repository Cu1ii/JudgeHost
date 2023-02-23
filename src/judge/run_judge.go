package judge

import "JudgeHost/src/models/dto"

func RunJudge(request *dto.JudgeRequest, res *string) error {
	return judge(request.SubmissionId,
		request.MemoryLimit,
		request.TimeLimit,
		request.ResolutionPath,
		request.SubmissionCode,
		request.Language,
		request.ProblemId,
		request.JudgePreference,
		request.Spj,
		res)
}
