package test

import (
	"fmt"
	"github.com/timandy/routine"
	"testing"
	"time"
)

var threadLocal = routine.NewThreadLocal()
var inheritableThreadLocal = routine.NewInheritableThreadLocal()

func TestThreadLocal(t *testing.T) {
	threadLocal.Set("hello world")
	inheritableThreadLocal.Set("Hello world2")
	fmt.Println("threadLocal:", threadLocal.Get())
	fmt.Println("inheritableThreadLocal:", inheritableThreadLocal.Get())

	// 子协程无法读取之前赋值的“hello world”。
	go func() {
		fmt.Println("threadLocal in goroutine:", threadLocal.Get())
		fmt.Println("inheritableThreadLocal in goroutine:", inheritableThreadLocal.Get())
	}()

	// 但是，可以通过 Go/GoWait/GoWaitResul 函数启动一个新的子协程，当前协程的所有可继承变量都可以自动传递。
	routine.Go(func() {
		fmt.Println("threadLocal in goroutine by Go:", threadLocal.Get())
		fmt.Println("inheritableThreadLocal in goroutine by Go:", inheritableThreadLocal.Get())
	})

	// 等待子协程执行完。
	time.Sleep(time.Second)
}
