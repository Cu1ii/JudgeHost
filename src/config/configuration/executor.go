package configuration

import (
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

/**
 * 判题线程池
 * 判题主要是cpu密集型的操作
 * 单核CPU上运行的多线程程序, 同一时间只能一个线程在跑
 * 系统帮你切换线程而已, 系统给每个线程分配时间片来执行
 * 每个时间片大概10ms左右, 看起来像是同时跑,
 * 但实际上是每个线程跑一点点就换到其它线程继续跑，
 * 多核并行量超过核心数目也有类似的道理
 * 但是，由于我们要计算判题时间，所以必须等于cpu核心数目
 * 否则这个值会出错（例如单核双并行，值就相差了两倍）
 * 这两个值根据用户容忍的等待时间以及测试时单机任务执行平均时长来获取自定义的判题线程池的相关配置
 */

var CORE_POOL_SIZE int = runtime.NumCPU()

const (
	EXPIRY_DURATION    = time.Second * 5
	PRE_ALLOC          = false
	MAX_BLOCKING_TASKS = 20
	NONBLOCKING        = true
)

var JudgeExecutorPool *ants.PoolWithFunc

func init() {
	var err error
	JudgeExecutorPool, err = ants.NewPoolWithFunc(CORE_POOL_SIZE, TaskFunc, ants.WithOptions(ants.Options{
		ExpiryDuration:   EXPIRY_DURATION,
		PreAlloc:         PRE_ALLOC,
		MaxBlockingTasks: MAX_BLOCKING_TASKS,
		Nonblocking:      NONBLOCKING,
	}))
	if err != nil {
		logrus.Error("judgeExecutorPool init fail: ", err.Error())
	}
}
