package queueworker

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

type workerFunc func(context.Context) error

type QueueWorker struct {
	dispatcherN int
	workerN     int
	workerFunc  workerFunc
	wg          *sync.WaitGroup
}

func New(dispatcherN int, workerN int, workerFunc workerFunc) *QueueWorker {
	return &QueueWorker{
		dispatcherN: dispatcherN,
		workerN:     workerN,
		workerFunc:  workerFunc,
		wg:          &sync.WaitGroup{},
	}
}

func (qw *QueueWorker) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-qw.listenCancelSignal(ctx)
		cancel()
	}()

	dispatcherChs := make([]<-chan interface{}, 0, qw.dispatcherN)
	for i := 0; i < qw.dispatcherN; i++ {
		dispatcherChs = append(dispatcherChs, qw.dispatch(ctx))
	}

	workerCh := qw.fanIn(ctx, dispatcherChs)
	for i := 0; i < qw.workerN; i++ {
		qw.work(ctx, workerCh)
	}

	qw.wg.Wait()
	log.Printf("Left behind goroutines: %d", runtime.NumGoroutine())
}

func (qw *QueueWorker) listenCancelSignal(ctx context.Context) <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGKILL)
	return ch
}

func (qw *QueueWorker) dispatch(ctx context.Context) <-chan interface{} {
	ch := make(chan interface{})

	qw.wg.Add(1)
	go func() {
		defer func() {
			close(ch)
			qw.wg.Done()
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			// 便宜的に 3秒ごとにメッセージを受け取る
			// SQSと連携する想定
			// output, err := sqs.ReceiveMessage(input)
			// msgs := output.Messages
			time.Sleep(time.Second * 3)
			msgs := []int{1, 2, 3, 4, 5}
			if len(msgs) == 0 {
				break
			}
			select {
			case <-ctx.Done():
				return
			default:
			}
			for _, m := range msgs {
				ch <- m
			}
		}
	}()

	return ch
}

func (qw *QueueWorker) fanIn(ctx context.Context, chs []<-chan interface{}) <-chan interface{} {
	multiplexedCh := make(chan interface{})

	multiplex := func(ch <-chan interface{}, ctx context.Context) {
		defer qw.wg.Done()

		for i := range ch {
			select {
			case <-ctx.Done():
				return
			case multiplexedCh <- i:
			}
		}
	}

	qw.wg.Add(len(chs))
	for _, ch := range chs {
		go multiplex(ch, ctx)
	}

	go func() {
		<-ctx.Done()
		close(multiplexedCh)
	}()

	return multiplexedCh
}

func (qw *QueueWorker) work(ctx context.Context, ch <-chan interface{}) {
	qw.wg.Add(1)
	go func() {
		defer qw.wg.Done()
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

		for msg := range ch {
			select {
			case <-ctx.Done():
				return
			default:
				if err := qw.workerFunc(context.WithValue(ctx, "message", msg)); err != nil {
					log.Printf("Error: %v", err)
				}
			}
		}
	}()
}
