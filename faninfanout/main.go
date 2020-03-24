package main

import (
	"context"
	"sync"
)

func fanIn(ctx context.Context, chs ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	multiplexedCh := make(chan interface{})

	multiplex := func(ch <-chan interface{}) {
		defer wg.Done()
		for i := range ch {
			select {
			case <-ctx.Done():
				return
			case multiplexedCh <- i:
			}
		}
	}

	wg.Add(len(chs))
	for _, ch := range chs {
		go multiplex(ch)
	}

	go func() {
		wg.Wait()
		close(multiplexedCh)
	}()

	return multiplexedCh
}
