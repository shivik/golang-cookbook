package basics

import "fmt"

// BasicTypes demonstrates Go's built-in scalar types.
func BasicTypes() {
	var i int = 42
	var u uint = 10
	var f float64 = 3.14
	var b bool = true
	var s string = "hello"
	var r rune = '界'
	var by byte = 0xFF

	fmt.Println("int:", i)
	fmt.Println("uint:", u)
	fmt.Println("float64:", f)
	fmt.Println("bool:", b)
	fmt.Println("string:", s)
	fmt.Println("rune (Unicode code point):", r, string(r))
	fmt.Println("byte (alias for uint8):", by)
}
