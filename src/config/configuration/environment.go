package configuration

import (
	"JudgeHost/src/common"
	"JudgeHost/src/util"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type JudgeEnvironmentConfiguration struct {
	JudgeEnvironment JudgeEnvironment `mapstructure:"judge-environment" json:"env" yaml:"judge-environment"`
}

type JudgeEnvironment struct {
	SubmissionPath string `mapstructure:"submission-path" yaml:"submission-path:"`
	ScriptPath     string `mapstructure:"script-path" yaml:"script-path:"`
	ResolutionPath string `mapstructure:"resolution-path" yaml:"resolution-path:"`
}

func (p *JudgeEnvironment) CheckJudgeEnvironmentBaseFileIn() (*common.UnifiedResponse, error) {
	res, err := util.IsDirectory(p.SubmissionPath)
	if err != nil {
		return common.NewUnifiedResponseMessgae("B1002"), err
	}
	if !res {
		return common.NewUnifiedResponseMessgae("B1002"), nil
	}

	res, err = util.IsDirectory(p.ResolutionPath)
	if err != nil {
		return common.NewUnifiedResponseMessgae("B1002"), err
	}
	if !res {
		return common.NewUnifiedResponseMessgae("B1002"), nil
	}

	res, err = util.IsDirectory(p.ScriptPath)
	if err != nil {
		return common.NewUnifiedResponseMessgae("B1002"), err
	}
	if !res {
		return common.NewUnifiedResponseMessgae("B1002"), nil
	}
	return nil, nil
}

var JudgeEnvironmentConfigurationEntity = JudgeEnvironmentInitializeConfig()

func JudgeEnvironmentInitializeConfig() *JudgeEnvironmentConfiguration {
	judgeEnvironmentConfigurationEntity := JudgeEnvironmentConfiguration{}
	configPath := "resources/config/judge-environment.yaml"
	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		logrus.Error("read config failed: ", err.Error())
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&judgeEnvironmentConfigurationEntity); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&judgeEnvironmentConfigurationEntity); err != nil {
		fmt.Println(err)
	}
	return &judgeEnvironmentConfigurationEntity
}

func GetJudgeEnvironmentConfiguration() *JudgeEnvironmentConfiguration {
	return JudgeEnvironmentConfigurationEntity
}
