package methods_interfaces

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// Area is a value receiver method.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Scale is a pointer receiver method that modifies the receiver.
func (r *Rectangle) Scale(f float64) {
	r.Width *= f
	r.Height *= f
}

func MethodsDemo() {
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println("area:", r.Area())
	r.Scale(2)
	fmt.Println("scaled area:", r.Area())
}
