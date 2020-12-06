package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zbcheng/xdays/user/internal/di"
)

func main() {
	container := di.BuildContainer()

	if err := container.Invoke(func(app *di.App) {
		fmt.Println("timer...")
		svr, err := grpc.New()
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
		for {
			s := <-c
			fmt.Printf("get a signal %s\n", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				// closeF
				fmt.Println("xdays exit")
				time.Sleep(time.Second)
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}); err != nil {
		fmt.Println(err)
	}
}
