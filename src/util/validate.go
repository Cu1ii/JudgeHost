package util

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	validate = validator.New()          // 实例化验证器
	chinese  = zh.New()                 // 获取中文翻译器
	uni      = ut.New(chinese, chinese) // 设置成中文翻译器
	trans, _ = uni.GetTranslator("zh")  // 获取翻译字典
)

func InitValidate() error {
	// 注册翻译器
	if err := zhs.RegisterDefaultTranslations(validate, trans); err != nil {
		logrus.Debug("register default translations fail: ", err.Error(), time.Now().UTC().String())
		return err
	}
	return nil
}

func ValidateStructCheck(verifyEntity interface{}) validator.ValidationErrorsTranslations {
	if err := validate.Struct(verifyEntity); err != nil {
		if verifyErrors, ok := err.(validator.ValidationErrors); ok {
			return verifyErrors.Translate(trans)
		}
		logrus.Debug("validate struct check error not validation errors: ", err.Error())
		return nil
	}
	return nil
}
