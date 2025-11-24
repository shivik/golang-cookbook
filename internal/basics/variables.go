package basics

import "fmt"

// Variables shows different ways to declare and initialize variables.
func Variables() {
	var explicit int
	explicit = 10

	var (
		a = 1
		b = 2
	)

	inferred := 3

	fmt.Println("explicit:", explicit)
	fmt.Println("a, b:", a, b)
	fmt.Println("inferred:", inferred)
}
