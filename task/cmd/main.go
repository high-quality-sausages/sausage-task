package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/high-quality-sausages/sausage-task/task/internal/di"
)

func main() {
	container := di.BuildContainer()

	if err := container.Invoke(func(app *di.App) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
		for {
			s := <-c
			fmt.Printf("get a signal %s\n", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				// closeF
				log.Println("[task] exit")
				time.Sleep(time.Second)
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}); err != nil {
		log.Fatalf("failed to start [task], err:%v\n", err)
	}
}
