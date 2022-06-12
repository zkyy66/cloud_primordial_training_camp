/**
 * @Date 2022/6/12
 * @Name grace-stop
 * @VariableName
**/
package module8

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GraceMain() {
	chanHandle := make(chan os.Signal)
	signal.Notify(chanHandle, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range chanHandle {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出程序。。。", s)
				graceExit()
			case syscall.SIGUSR1:
				fmt.Println("user1 signal..", s)
			case syscall.SIGUSR2:
				fmt.Println("user2 signal...", s)
			default:
				fmt.Println("其他。。。。", s)
			}
		}
	}()
	fmt.Println("开始。。。")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(2 * time.Second)
	}
}
func graceExit() {
	fmt.Println("准备退出。。。")
	fmt.Println("执行。。。。")
	fmt.Println("已经退出。。。。")
	os.Exit(0)
}
