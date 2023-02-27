package util

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func StringListToString(stringList []string) string {
	result := ""
	for _, str := range stringList {
		result += str
	}
	return result
}

func TransToJson(obj interface{}) (string, error) {
	marshal, err := json.Marshal(obj)
	if err != nil {
		logrus.Error("transfer responseVo to json fail in compileError:16 ", err)
		return "system error " + err.Error(), err
	}
	return string(marshal), nil
}

func TransformResultToString(result int64) string {
	if result == 2 || result == 1 {
		return "Time Limit Exceeded"
	}
	if result == 3 {
		return "Memory Limit Exceeded"
	}
	if result == 4 {
		return "Runtime Error"
	}
	if result == 5 {
		return "System Error"
	}
	if result == -5 {
		return "Presentation Error"
	}
	if result == -3 {
		return "Wrong Answer"
	}
	return "Unknown Answer"
}
