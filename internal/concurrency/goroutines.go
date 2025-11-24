package concurrency

import (
	"fmt"
	"time"
)

// GoroutinesDemo shows basic goroutine usage.
func GoroutinesDemo() {
	fmt.Println("start GoroutinesDemo")

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("goroutine 1:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("goroutine 2:", i)
			time.Sleep(150 * time.Millisecond)
		}
	}()

	time.Sleep(600 * time.Millisecond)
	fmt.Println("end GoroutinesDemo")
}
