package vo

import "JudgeHost/src/models/entity"

type ResponseVo struct {
	CpuTime  int                 `json:"cpu_time"`
	Memory   int                 `json:"memory"`
	Result   int                 `json:"result"`
	Msg      string              `json:"msg"` // 用来记录 result 对应的判题结果 如果是编译错误则记录编译信息
	CaseList []entity.CaseStatus `json:"case_list"`
	Testcase string              `json:"testcase"` // 记录最后一个测评的样例
	Error    int                 `json:"error"`    // judger 错误码 建议在 result 显示出现 re | se 的时候再查看
}

func (r *ResponseVo) AppendCase(caseStatus *entity.CaseStatus) {
	r.CaseList = append(r.CaseList, *caseStatus)
}
