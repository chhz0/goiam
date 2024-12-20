package ginserver

import (
	"os"
	"os/signal"
)

var (
	onlyOneSignalHandler = make(chan struct{})
	shutdownHandler      chan os.Signal
)

// SetUpSignalHandler 注册信号处理函数, 返回一个停止通道
// 当收到信号时, 会关闭停止通道， 如果捕捉到第二个信号，以退出码 1 退出程序
func SetUpSignalHandler() <-chan struct{} {
	close(onlyOneSignalHandler)

	signal.Notify(shutdownHandler, shutdownSignal...)

	quit := make(chan struct{})

	go func() {
		<-shutdownHandler
		close(quit)
		<-shutdownHandler
		os.Exit(1)
	}()

	return quit
}
