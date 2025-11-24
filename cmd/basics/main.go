package main

import (
	"fmt"

	"github.com/shivik/go-lang-cookbook/internal/basics"
)

// This command demonstrates Go's basic types, variables, constants,
// arrays, slices, maps, and structs.
func main() {
	fmt.Println("=== BASIC TYPES ===")
	basics.BasicTypes()
	fmt.Println()

	fmt.Println("=== VARIABLES ===")
	basics.Variables()
	fmt.Println()

	fmt.Println("=== CONSTANTS ===")
	basics.Constants()
	fmt.Println()

	fmt.Println("=== ARRAYS, SLICES, MAPS ===")
	basics.ArraysSlicesMaps()
	fmt.Println()

	fmt.Println("=== STRUCTS ===")
	basics.Structs()
}
