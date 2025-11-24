package functions

import "fmt"

// add demonstrates a simple function with multiple return values.
func add(a, b int) (sum int) {
	sum = a + b
	return
}

// divide returns a result and an error using Go's error handling style.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// FunctionsDemo shows function literals, closures, and defer.
func FunctionsDemo() {
	fmt.Println("add(2, 3):", add(2, 3))

	result, err := divide(10, 2)
	fmt.Println("divide(10, 2):", result, "err:", err)

	_, err = divide(5, 0)
	fmt.Println("divide(5, 0) err:", err)

	// Anonymous function and closure.
	counter := 0
	increment := func() int {
		counter++
		return counter
	}
	fmt.Println("closure:", increment(), increment(), increment())

	// defer example: LIFO execution when surrounding function returns.
	fmt.Println("defer demo start")
	defer fmt.Println("deferred 1 (runs last)")
	defer fmt.Println("deferred 2")
	fmt.Println("defer demo end")
}
