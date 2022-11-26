package bo

import (
	global "JudgeHost/src/config/global"
	"JudgeHost/src/models/dto"
)

// JudgeConfigurationBO 判题配置，可以理解为全局变量
type JudgeConfigurationBO struct {
	SubmissionPath string
	SubmissionId   int64
	ExtraInfo      []string
	JudgeConfig    *dto.JudgeDTO
	Runner         string
	WorkPath       string
	ScriptPath     string
	ResolutionPath string
}

func NewJudgeConfigurationBO(judgeConfig *dto.JudgeDTO,
	workPath, srciptPath, resolutionPath string) *JudgeConfigurationBO {
	global.GlobalSubmissionId.Add(1)
	return &JudgeConfigurationBO{
		SubmissionId:   global.GlobalSubmissionId.Load(),
		JudgeConfig:    judgeConfig,
		WorkPath:       workPath,
		ScriptPath:     srciptPath,
		ResolutionPath: resolutionPath,
	}
}
