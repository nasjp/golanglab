package workdispacher

import (
	"context"
	"log"
	"sync"
)

type Dispatcher struct {
	sem chan struct{}
	wg  sync.WaitGroup
}

type WorkFunc func(context.Context) error

func NewDispatcher(max int) *Dispatcher {
	return &Dispatcher{
		sem: make(chan struct{}, max),
	}
}

func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

func (d *Dispatcher) Work(ctx context.Context, proc WorkFunc) {
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		if err := d.work(ctx, proc); err != nil {
			log.Printf("Error: %v", err)
		}
	}()
}

func (d *Dispatcher) work(ctx context.Context, proc WorkFunc) error {
	select {
	case <-ctx.Done():
		log.Printf("cancel work")
		return nil
	case d.sem <- struct{}{}:
		// got semaphore
		defer func() { <-d.sem }()
	}

	return proc(ctx)
}
