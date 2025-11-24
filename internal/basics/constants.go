package basics

import "fmt"

// Constants shows typed and untyped constants and iota.
func Constants() {
	const pi = 3.14159
	const typedPi float64 = 3.14159

	const (
		_ = iota
		StatusPending
		StatusRunning
		StatusDone
	)

	fmt.Println("pi:", pi)
	fmt.Println("typedPi:", typedPi)
	fmt.Println("StatusPending, StatusRunning, StatusDone:",
		StatusPending, StatusRunning, StatusDone)
}
