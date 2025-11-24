package concurrency

import (
	"fmt"
	"time"
)

// SelectDemo demonstrates select over multiple channels.
func SelectDemo() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		c1 <- "from c1"
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		c2 <- "from c2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println("received:", m1)
		case m2 := <-c2:
			fmt.Println("received:", m2)
		}
	}
}
