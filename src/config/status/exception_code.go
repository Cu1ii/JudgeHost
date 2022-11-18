package status

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Configuration struct {
	Exceptions JudgeHostExceptions `mapstructure:"judge-host-exceptions" json:"judge-host-exceptions" yaml:"judge-host-exceptions"`
}

type JudgeHostExceptions struct {
	ExceptionCodes map[string]string `mapstructure:"codes" json:"codes" yaml:"codes"`
}

var StatusApp *Configuration = StatusInitializeConfig()

func StatusInitializeConfig() *Configuration {
	statusApp := Configuration{}
	configPath := "resources/config/exception-codes.yaml"
	StatusApp.Exceptions.ExceptionCodes = make(map[string]string)
	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&statusApp); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&statusApp); err != nil {
		fmt.Println(err)
	}
	return &statusApp
}

func GetConfiguration() *Configuration {
	return StatusApp
}
