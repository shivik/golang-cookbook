package generics

import "fmt"

// SumNumbers uses a type constraint supporting both ints and floats.
type Number interface {
	~int | ~int64 | ~float64
}

func SumNumbers[T Number](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

// Map applies a function to each element of a slice.
func Map[T any, U any](in []T, fn func(T) U) []U {
	out := make([]U, 0, len(in))
	for _, v := range in {
		out = append(out, fn(v))
	}
	return out
}

func GenericsDemo() {
	ints := []int{1, 2, 3}
	floats := []float64{1.5, 2.5, 3.0}

	fmt.Println("SumNumbers(ints):", SumNumbers(ints))
	fmt.Println("SumNumbers(floats):", SumNumbers(floats))

	names := []string{"go", "lang"}
	lengths := Map(names, func(s string) int { return len(s) })
	fmt.Println("Map lengths:", lengths)
}
