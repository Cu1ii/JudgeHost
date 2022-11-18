package main

import (
	"JudgeHost/src/config/configuration"
	"JudgeHost/src/config/global"
	"JudgeHost/src/controllers"
	"JudgeHost/src/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	r.Use(log.LogMiddleWare())
	controllers.LoadControllers(r)
	if err := r.Run(":" + strconv.FormatInt(int64(global.GlobalApp.App.Port), 10)); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
	defer configuration.JudgeExecutorPool.Release()

}
