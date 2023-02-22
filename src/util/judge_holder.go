package util

import (
	"JudgeHost/src/models/bo"
	"JudgeHost/src/models/dto"
	"github.com/timandy/routine"
	"strconv"
)

const (
	CompileScriptName   = "compile.sh"
	JudgeCoreScriptName = "y_judge"
	CompareScriptName   = "compare.sh"
	CodeFileName        = "Main"
	CompileStdOutName   = "compile.out"
	CompileStdErrName   = "compile.err"
	RunnerScriptName    = "run"
	SpjScriptName       = "spj.sh"
	CheckerName         = "checker"
)

var JudgeHolderThreadLocal = routine.NewThreadLocal()

func InitJudgeConfiguration(judgeConfigurationBO *bo.JudgeConfigurationBO) {
	JudgeHolderThreadLocal.Set(judgeConfigurationBO)
}

func RemoveThreadLocal() {
	JudgeHolderThreadLocal.Remove()
}

func GetJudgeConfiguration() *bo.JudgeConfigurationBO {
	return JudgeHolderThreadLocal.Get().(*bo.JudgeConfigurationBO)
}

func GetCompileScriptPath() string {
	return GetJudgeConfiguration().ScriptPath + "/" + CompileScriptName
}

func GetJudgeCoreScriptPath() string {
	return GetJudgeConfiguration().ScriptPath + "/" + JudgeCoreScriptName
}

func GetCompareScriptPath() string {
	return GetJudgeConfiguration().ScriptPath + "/" + CompareScriptName
}

func GetSpjScriptPath() string {
	return GetJudgeConfiguration().ScriptPath + "/" + SpjScriptName
}

func GetRunnerScriptPath() string {
	return GetSubmissionWorkingPath() + "/" + RunnerScriptName
}

func GetJudgeConfig() *dto.JudgeRequest {
	return GetJudgeConfiguration().JudgeConfig
}

func GetCodePath(extension string) string {
	return GetSubmissionWorkingPath() + "/" + CodeFileName + "." + extension
}

func GetSubmissionId() string {
	return strconv.FormatInt(GetJudgeConfiguration().SubmissionId, 10)
}

func GetRunner() string {
	return GetJudgeConfiguration().Runner
}

func GetSubmissionWorkingPath() string {
	return GetJudgeConfiguration().WorkPath + "/" + GetSubmissionId()
}

func GetResolutionPath() string {
	return GetJudgeConfiguration().ResolutionPath
}

func GetExtraInfo() []string {
	var extraInfo []string
	copy(extraInfo, GetJudgeConfiguration().ExtraInfo)
	return extraInfo
}

func SetExtraInfo(extraInfo []string) {
	configurationBO := GetJudgeConfiguration()
	configurationBO.ExtraInfo = extraInfo
}
