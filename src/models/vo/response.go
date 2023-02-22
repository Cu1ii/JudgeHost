package vo

import "JudgeHost/src/models/entity"

type ResponseVo struct {
	RealTime int                 `json:"real_time"`
	CpuTime  int                 `json:"cpu_time"`
	Memory   int                 `json:"memory"`
	Result   int                 `json:"result"`
	Msg      string              `json:"msg"` // 用来记录 result 对应的判题结果 如果是编译错误则记录编译信息
	CaseList []entity.CaseStatus `json:"case_list"`
	Error    int                 `json:"error"`
	Signal   int                 `json:"signal"`
	ExitCode int                 `json:"exit_code"`
}

func NewResponseVo(result *entity.Result) *ResponseVo {
	return &ResponseVo{
		RealTime: result.RealTime,
		CpuTime:  result.CpuTime,
		Memory:   result.Memory,
		Result:   result.Result,
		CaseList: make([]entity.CaseStatus, 0),
		Error:    result.Error,
		Signal:   result.Signal,
		ExitCode: result.ExitCode,
	}
}

func (r *ResponseVo) AppendCase(caseStatus *entity.CaseStatus) {
	r.CaseList = append(r.CaseList, *caseStatus)
}
