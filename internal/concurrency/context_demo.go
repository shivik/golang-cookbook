package concurrency

import (
	"context"
	"fmt"
	"time"
)

// ContextDemo shows graceful cancellation with context.
func ContextDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("worker cancelled:", ctx.Err())
				return
			case ch <- i:
				time.Sleep(80 * time.Millisecond)
			}
		}
	}()

	for v := range ch {
		fmt.Println("got value:", v)
	}
}
