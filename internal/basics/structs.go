package basics

import "fmt"

// Person demonstrates struct definition and initialization.
type Person struct {
	Name string
	Age  int
}

// NewPerson is a simple constructor function.
func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

// Structs demonstrates struct usage and pointers to structs.
func Structs() {
	p := NewPerson("Gopher", 10)
	fmt.Println("person value:", p)

	pp := &p
	pp.Age++
	fmt.Println("person via pointer:", *pp)
}
