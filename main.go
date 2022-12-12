package main

import "JudgeHost/xoj"

func main() {
	//r := gin.Default()
	//r.Use(logs.LogMiddleWare())
	//controllers.LoadControllers(r)
	//if err := r.Run(":" + strconv.FormatInt(int64(global.GlobalApp.App.Port), 10)); err != nil {
	//	fmt.Println("startup service failed, err:%v\n", err)
	//}
	//defer configuration.JudgeExecutorPool.Release()
	xoj.RunJudge()
}
