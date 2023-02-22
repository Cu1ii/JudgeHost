package judge

import (
	"JudgeHost/src/models/vo"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func compileError(msg string) string {
	res := vo.ResponseVo{
		Result: -4,
		Msg:    msg,
	}
	marshal, err := json.Marshal(res)
	if err != nil {
		logrus.Error("transfer responseVo to json fail in compileError:16 ", err)
		return "system error"
	}
	return string(marshal)
}
