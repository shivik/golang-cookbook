package main

import (
	"fmt"

	"github.com/shivik/go-lang-cookbook/internal/methods_interfaces"
)

func main() {
	fmt.Println("=== METHODS ===")
	methods_interfaces.MethodsDemo()
	fmt.Println()

	fmt.Println("=== INTERFACES ===")
	methods_interfaces.InterfacesDemo()
	fmt.Println()

	fmt.Println("=== EMBEDDING ===")
	methods_interfaces.EmbeddingDemo()
}
