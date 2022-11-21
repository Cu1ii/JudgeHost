package global

import (
	"JudgeHost/src/config/app"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

// 全局配置

var GlobalSubmissionId int64

var JudgeConfigDefault = map[string]int{
	// cpu时间限制 4000ms
	"TIME_LIMIT_DEFAULT": 4000,
	// 内存限制 1024 * 32 kb = 32mb
	"MEMORY_LIMIT_DEFAULT": 1024 * 32,
	// 实际时间限制 4000ms
	"WALL_TIME_DEFAULT": 4000,
	// 进程限制
	"PROCESS_LIMIT_DEFAULT": 1,
	// 输出限制
	"OUTPUT_LIMIT_DEFAULT": 1000000,
}

var JudgePreference = map[string]string{
	"ACM": "ACM",
	"OI":  "OI",
}

func ToJudgePreference(preference string) (string, error) {
	if _, ok := JudgePreference[preference]; ok {
		return JudgePreference[preference], nil
	}
	return "", errors.New("preference not exist")
}

var JudgeResult = map[string]int{
	"ACCEPT":               0,
	"WRONG_ANSWER":         1,
	"RUNTIME_ERROR":        2,
	"TIME_LIMIT_EXCEEDED":  3,
	"MEMORY_LIMIT_EXCEED":  4,
	"OUTPUT_LIMIT_EXCEED":  5,
	"SEGMENTATION_FAULT":   6,
	"FLOAT_ERROR":          7,
	"UNKNOWN_ERROR":        8,
	"INPUT_FILE_NOT_FOUND": 9,
	"CAN_NOT_MAKE_OUTPUT":  10,
	"SET_LIMIT_ERROR":      11,
	"NOT_ROOT_USER":        12,
	"FORK_ERROR":           13,
	"CREATE_THREAD_ERROR":  14,
	"VALIDATE_ERROR":       15,
	"COMPILE_ERROR":        16,
}

var ReverseJudgeResult = []string{"ACCEPT", "WRONG_ANSWER", "RUNTIME_ERROR", "TIME_LIMIT_EXCEEDED",
	"MEMORY_LIMIT_EXCEED", "OUTPUT_LIMIT_EXCEED", "SEGMENTATION_FAULT", "FLOAT_ERROR",
	"UNKNOWN_ERROR", "INPUT_FILE_NOT_FOUND", "CAN_NOT_MAKE_OUTPUT", "SET_LIMIT_ERROR",
	"NOT_ROOT_USER", "FORK_ERROR", "CREATE_THREAD_ERROR", "VALIDATE_ERROR", "COMPILE_ERROR",
}

func ToJudgeResultType(number int) (res string, err error) {
	if number >= len(ReverseJudgeResult) {
		return "", errors.New("out of range")
	}
	return ReverseJudgeResult[number], nil
}

var LanguageScript = map[string]string{
	"PYTHON": "#!/bin/sh\n" +
		"\n" +
		"echo '#!/bin/sh' > run" +
		"\n" +
		"echo 'python3 CODE_PATH' >> run",
	"JAVA": "#!/bin/sh\n" +
		"\n" +
		"javac CODE_PATH" +
		"\n" +
		"echo '#!/bin/sh' > run" +
		"\n" +
		"echo 'cd SUBMISSION_PATH' >> run" +
		"\n" +
		"echo 'java Main' >> run",
	"C": "#!/bin/sh\n" +
		"\n" +
		"gcc -Wall -O2 -std=gnu11 CODE_PATH -o run -lm",
	"C_PLUS_PLUS": "#!/bin/sh\n" +
		"\n" +
		"g++ -Wall -O2 CODE_PATH -o run",
}

var Language = map[string]string{
	"PYTHON":      "py",
	"JAVA":        "java",
	"C":           "c",
	"C_PLUS_PLUS": "cpp",
}

func GetBuildScriptByRunningPath(languageType, submissionPath, codePath string) (path string, err error) {
	if script, ok := LanguageScript[languageType]; ok {
		path = strings.Replace(strings.Replace(script, "SUBMISSION_PATH", submissionPath, -1),
			"CODE_PATH", codePath, -1)
		return
	}
	return "", errors.New("languageScript not exist")
}

var GlobalApp *app.Configuration = appInitializeConfig()

func appInitializeConfig() *app.Configuration {
	globeApplication := app.Configuration{}
	configPath := "resources/application.yaml"
	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("read config failed: %s \n", err)
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&globeApplication); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&globeApplication); err != nil {
		fmt.Println(err)
	}
	return &globeApplication
}

func GetGlobalConfiguration() *app.Configuration {
	return GlobalApp
}
