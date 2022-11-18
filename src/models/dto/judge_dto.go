package dto

type JudgeDTO struct {
	SubmissionCode  string `form:"submissionCode" validate:"required"`        // 代码不得为空
	RealTimeLimit   int    `form:"realTimeLimit" validate:"lte=10000"`        // 实际时间限制最大为10 * 1000ms
	CpuTimeLimit    int    `form:"cpuTimeLimit" validate:"lte=10000"`         // cpu 时间限制最大为10 * 1000ms
	MemoryLimit     int    `form:"memoryLimit" validate:"gte=3000,lte=65536"` // 内存限制最小为 3000kb, 最大限制为 65536kb
	OutputLimit     int    `form:"outputLimit" validate:"gte=10"`             // 输出最小限制为 10Byte
	Language        string `form:"language" validate:"required"`              // 语言不得为空
	JudgePreference string
	Solutions       []*SolutionDTO `form:"solutions" validate:"lte=10,gte=0"` // 期望输入, 输出长度最小为 1 最大为 10
}

func (p *JudgeDTO) IsAcmMode() bool {
	if p.JudgePreference == "ACM" {
		return true
	}
	return false
}
