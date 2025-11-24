package main

import (
	"fmt"

	"github.com/shivik/go-lang-cookbook/internal/concurrency"
)

func main() {
	fmt.Println("=== GOROUTINES ===")
	concurrency.GoroutinesDemo()
	fmt.Println()

	fmt.Println("=== CHANNELS ===")
	concurrency.ChannelsDemo()
	fmt.Println()

	fmt.Println("=== SELECT ===")
	concurrency.SelectDemo()
	fmt.Println()

	fmt.Println("=== CONTEXT CANCELLATION ===")
	concurrency.ContextDemo()
}
