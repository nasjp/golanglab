package main

import "fmt"

func main() {
	closeCh()
}

func closeCh() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 1
	fmt.Println(<-ch)
	close(ch)
	// クローズしても入ってる値は取れる
	fmt.Println(<-ch)
	// なくなったらゼロ値が返る
	fmt.Println(<-ch)
}

func rangeCh() {
	ch := owner()
	for i := range ch {
		fmt.Println(i)
	}
}

func owner() <-chan int {
	length := 4
	ch := make(chan int, length)
	go func() {
		defer close(ch)
		for i := 0; i < length; i++ {
			ch <- i
		}
	}()
	return ch
}
