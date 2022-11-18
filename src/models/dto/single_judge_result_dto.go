package dto

import (
	"JudgeHost/src/config/global"
	"fmt"
)

type SingleJudgeResultDTO struct {
	RealTimeCost string
	MemoryCost   string
	CpuTimeCost  string
	Condition    int
	StdInPath    string
	StdOutPath   string
	StdErrPath   string
	Message      string
}

func (p *SingleJudgeResultDTO) SetMessage() {
	var err error
	p.Message, err = global.ToJudgeResultType(p.Condition)
	if err != nil {
		fmt.Println(err)
	}
}
