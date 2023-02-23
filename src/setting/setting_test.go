package setting

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"testing"
)

var (
	app              = App{}
	judgeEnvironment = JudgeEnvironment{}
	exceptionCodes   = JudgeHostExceptions{}
)

func TestSetting_ReadSection(t *testing.T) {

	err := setupSetting()
	if err != nil {
		logrus.Fatalf("init.setupSetting err: %v", err)
	}

	fmt.Println(app)
	fmt.Println(judgeEnvironment)
	fmt.Println(exceptionCodes)
}

func setupSetting() error {
	vp := viper.New()
	vp.SetConfigName("application")
	vp.AddConfigPath("/home/cu1/Project/Go/JudgeHost/resources/config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return err
	}
	set, err := &Setting{vp}, nil
	if err != nil {
		return err
	}

	err = set.ReadSection("app", &app)
	if err != nil {
		return err
	}

	err = set.ReadSection("judge-environment", &judgeEnvironment)
	if err != nil {
		return err
	}

	err = set.ReadSection("judge-host-exceptions", &exceptionCodes)
	if err != nil {
		return err
	}

	return nil
}
