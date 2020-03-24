package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nasjp/golanglab/workdispatcher/services"
	"github.com/nasjp/golanglab/workdispatcher/workdispacher"
)

const (
	workerN     = 10
	dispatcherN = 3
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)

	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		<-sigCh
		cancel()
	}()

	d := workdispacher.NewDispatcher(dispatcherN)
	for i := 0; i < workerN; i++ {
		ctx = context.WithValue(ctx, "id", int64(i))
		d.Work(ctx, services.SampleHandler)
	}

	d.Wait()
}
