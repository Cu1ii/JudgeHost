package dto

type JudgeRequest struct {
	ProblemId       int    `form:"problemId" validate:"required" json:"problem_id"`
	SubmissionId    int    `form:"submissionId" validate:"required" json:"submission_id"`
	SubmissionCode  string `form:"submissionCode" validate:"required" json:"submission_code,omitempty"` // 代码不得为空
	ResolutionPath  string `form:"resolutionPath" json:"resolution_path"`
	TimeLimit       int    `form:"cpuTimeLimit" validate:"lte=10000" json:"cpu_time_limit,omitempty"`       // cpu 时间限制最大为10 * 1000ms
	MemoryLimit     int    `form:"memoryLimit" validate:"gte=3000,lte=65536" json:"memory_limit,omitempty"` // 内存限制最小为 3000kb, 最大限制为 65536kb
	OutputLimit     int    `form:"outputLimit" validate:"gte=10" json:"output_limit,omitempty"`             // 输出最小限制为 10Byte
	Language        string `form:"language" validate:"required" json:"language,omitempty"`                  // 语言不得为空
	JudgePreference int    `form:"judgePreference" json:"judge_preference"`                                 // 判题类型 acm = 0 / oi = 1 默认为 acm 模式
	Spj             bool   `json:"spj"`                                                                     // 判题模式 非 spj = false / spj = true默认为非 spj
}

func (p *JudgeRequest) IsAcmMode() bool {
	if p.JudgePreference == 0 {
		return true
	}
	return false
}

func (p *JudgeRequest) IsSpj() bool {
	return p.Spj
}
