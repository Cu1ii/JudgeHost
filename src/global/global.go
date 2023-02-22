package global

import (
	"JudgeHost/src/setting"
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

const (
	WAITING                  = -6
	PRESENTATION_ERROR       = -5
	COMPILE_ERROR            = -4
	WRONG_ANSWER             = -3
	PENDING                  = -1
	JUDGINNG                 = -2
	CPU_TIME_LIMIT_EXCEEDED  = 1
	REAL_TIME_LIMIT_EXCEEDED = 2
	MEMORY_LIMIT_EXCEEDED    = 3
	RUNTIME_ERROR            = 4
	SYSTEM_ERROR             = 5
)

var (
	Logger                  = (*logrus.Logger)(nil)
	AppSetting              = (*setting.App)(nil)
	JudgeEnvironmentSetting = (*setting.JudgeEnvironment)(nil)
	ExceptionCodes          = (*setting.JudgeHostExceptions)(nil)
	JudgeExecutorPool       = (*ants.Pool)(nil)
)
