package bo

import (
	"JudgeHost/src/models/dto"
)

// JudgeConfigurationBO 判题配置，可以理解为全局变量
type JudgeConfigurationBO struct {
	SubmissionPath string
	SubmissionId   int64
	ExtraInfo      []string
	JudgeConfig    *dto.JudgeRequest
	Runner         string
	WorkPath       string
	ScriptPath     string
	ResolutionPath string
}

func NewJudgeConfigurationBO(judgeConfig *dto.JudgeRequest,
	workPath, srciptPath, resolutionPath string) *JudgeConfigurationBO {
	//global.GlobalSubmissionId.Add(1)
	return &JudgeConfigurationBO{
		SubmissionId:   1,
		JudgeConfig:    judgeConfig,
		WorkPath:       workPath,
		ScriptPath:     srciptPath,
		ResolutionPath: resolutionPath,
	}
}
