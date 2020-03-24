package queueworker

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type workerFunc func(context.Context) error

type QueueWorker struct {
	dispatcherN int
	workerN     int
	workerFunc  workerFunc
}

func New(dispatcherN int, workerN int, workerFunc workerFunc) *QueueWorker {
	return &QueueWorker{
		dispatcherN: dispatcherN,
		workerN:     workerN,
		workerFunc:  workerFunc,
	}
}

func (qw *QueueWorker) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	qw.listenCancelSignal(ctx)

	dispatcherChs := make([]<-chan interface{}, 0, qw.dispatcherN)
	for i := 0; i < qw.dispatcherN; i++ {
		dispatcherChs = append(dispatcherChs, qw.dispatch(ctx))
	}

	workerCh := qw.fanIn(ctx, dispatcherChs...)
	for i := 0; i < qw.workerN; i++ {
		qw.work(ctx, workerCh)
	}

	<-ctx.Done()
}

func (qw *QueueWorker) listenCancelSignal(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-sigCh
		close(sigCh)
		cancel()
	}()
}

func (qw *QueueWorker) fanIn(ctx context.Context, chs ...<-chan interface{}) <-chan interface{} {
	multiplexedCh := make(chan interface{})

	multiplex := func(ch <-chan interface{}) {
		for i := range ch {
			select {
			case <-ctx.Done():
				return
			case multiplexedCh <- i:
			}
		}
	}

	for _, ch := range chs {
		go multiplex(ch)
	}

	go func() {
		<-ctx.Done()
		close(multiplexedCh)
	}()

	return multiplexedCh
}

func (qw *QueueWorker) dispatch(ctx context.Context) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// 便宜的に 3秒ごとにメッセージを受け取る
				// SQSと連携する想定
				// output, err := sqs.ReceiveMessage(input)
				// msgs := output.Messages
				time.Sleep(time.Second * 3)
				msgs := []int{1, 2, 3, 4, 5}
				if len(msgs) == 0 {
					break
				}
				for _, m := range msgs {
					ch <- m
				}
			}
		}
	}()

	go func() {
		<-ctx.Done()
		close(ch)
	}()
	return ch
}

func (qw *QueueWorker) work(ctx context.Context, ch <-chan interface{}) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover: %v", err)
				select {
				case <-ctx.Done():
					return
				default:
					qw.work(ctx, ch)
				}
			}
		}()
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-ch:
				fmt.Println("goroutines: ", runtime.NumGoroutine())
				ctx = context.WithValue(ctx, "message", msg)
				err := qw.workerFunc(ctx)
				if err != nil {
					log.Printf("Error: %v", err)
					// time.Sleep(10 * time.Second)
				}
			}
		}
	}()
}
