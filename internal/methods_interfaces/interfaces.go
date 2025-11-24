package methods_interfaces

import "fmt"

// Shape is an interface type.
type Shape interface {
	Area() float64
}

// DescribeShape shows dynamic dispatch via interfaces.
func DescribeShape(s Shape) {
	fmt.Printf("shape area: %.2f\n", s.Area())
}

func InterfacesDemo() {
	r := Rectangle{Width: 5, Height: 6}
	DescribeShape(r)

	var s Shape = r
	fmt.Printf("stored in interface: %#v\n", s)
}
