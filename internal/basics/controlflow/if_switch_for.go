package controlflow

import "fmt"

// IfSwitch demonstrates if and switch statements.
func IfSwitch() {
	x := 7

	if x%2 == 0 {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "is odd")
	}

	switch {
	case x < 0:
		fmt.Println("negative")
	case x == 0:
		fmt.Println("zero")
	default:
		fmt.Println("positive")
	}

	day := 3
	switch day {
	case 1, 2, 3:
		fmt.Println("early in the week")
	case 4, 5:
		fmt.Println("late in the week")
	default:
		fmt.Println("weekend")
	}
}

// ForLoop demonstrates for loops, including range.
func ForLoop() {
	for i := 0; i < 3; i++ {
		fmt.Println("classic for:", i)
	}

	n := 3
	for n > 0 {
		fmt.Println("while-style:", n)
		n--
	}

	names := []string{"go", "rust", "python"}
	for idx, name := range names {
		fmt.Printf("index=%d value=%s\n", idx, name)
	}
}
