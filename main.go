package main

import (
	"JudgeHost/src/global"
	"JudgeHost/src/judge"
	"JudgeHost/src/logs"
	"JudgeHost/src/setting"
	"JudgeHost/src/util"
	"JudgeHost/src/util/pool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func init() {
	logs.InitLog()

	if err := setupSetting(); err != nil {
		logrus.Fatalf("init.setupSetting err: %v", err)
	}

	if err := pool.InitPool(); err != nil {
		logrus.Fatalf("init.pool err: %v", err)
	}

	if err := util.InitValidate(); err != nil {
		logrus.Fatalf("init.validate err: %v", err)
	}

}

func setupSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = set.ReadSection("app", &global.AppSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("judge-environment", &global.JudgeEnvironmentSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("judge-host-exceptions", &global.ExceptionCodes)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(judge.JudgeServiceImpl))
	lis, err := net.Listen("tcp", ":"+strconv.FormatInt(int64(global.AppSetting.Port), 10))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}

//defer global.JudgeExecutorPool.Release()
//r := gin.Default()
//r.Use(middleware.LogMiddleWare(global.Logger))
//controllers.LoadControllers(r)
//if err := r.Run(":" + strconv.FormatInt(int64(global.AppSetting.Port), 10)); err != nil {
//fmt.Println("startup service failed, err:%v\n", err)
//}
