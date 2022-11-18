package service

import "JudgeHost/src/config/configuration"

type JudgeService struct {
	JudgeEnvironmentConfiguration *configuration.JudgeEnvironmentConfiguration
	SolutionStdInPathKey          string
	SolutionExpectedStdOutPathKey string
	EnableJudgeCoreGuard          int
	DisableJudgeCoreGuard         int
	UseRootUid                    int
	UseDefaultUid                 int
	CompileOutMaxSize             int
}

func NewJudgeService(environmentConfiguration *configuration.JudgeEnvironmentConfiguration) *JudgeService {
	return &JudgeService{
		JudgeEnvironmentConfiguration: environmentConfiguration,
		SolutionStdInPathKey:          "stdIn",
		SolutionExpectedStdOutPathKey: "expectedStdOut",
		EnableJudgeCoreGuard:          1,
		DisableJudgeCoreGuard:         0,
		UseRootUid:                    0,
		UseDefaultUid:                 6666,
		CompileOutMaxSize:             100000,
	}
}

// CompileSubmission 读取 compile.sh 生成脚本
func (s *JudgeService) CompileSubmission() []string {
	return nil
}
