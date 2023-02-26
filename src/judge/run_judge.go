package judge

func RunJudge(request *JudgeRequest) (*JudgeResponse, error) {
	return judge(request.SubmissionId,
		request.MemoryLimit,
		request.TimeLimit,
		request.ResolutionPath,
		request.SubmissionCode,
		request.Language,
		request.ProblemId,
		request.JudgePreference,
		request.Spj,
	)
}
