package concurrency

import (
	"fmt"
	"time"
)

// ChannelsDemo shows unbuffered and buffered channels.
func ChannelsDemo() {
	ch := make(chan string)

	go func() {
		ch <- "hello from goroutine"
	}()

	msg := <-ch
	fmt.Println("received:", msg)

	buf := make(chan int, 2)
	buf <- 1
	buf <- 2
	close(buf)

	for v := range buf {
		fmt.Println("buffered value:", v)
	}

	// Small timeout pattern using select and a timer.
	done := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("done signal received")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("timeout")
	}
}
