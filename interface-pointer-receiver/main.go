package main

import(
	"fmt"
)

type Rect struct {
	width, height float64
}
type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r *Rect) Area() float64 {
	(*r).width = 3	// since pointer is passed, we can update the original
	r.height = 3
	return r.width * r.height
	// remember in Go, struct_ptr.property works like struct.property
}
func (r Rect) Perimeter() float64 {
	r.width = 5		// this doesn't change the original as a copied value was passed
	r.height = 5
	return 2 * (r.width + r.height)
	// but this is a gotchas. It looks Go pass a copy of the value to the Perimeter method
}

func main() {
	r := Rect{2.0, 2.0}
	//var s Shape = r	// A copied value of r will be pass to the implementation methods of Shape
	var s Shape = &r	// An addr of r (pointer) will be passed to the implementation methods of Shape

	fmt.Println("r: ", r)
	area := s.Area()
	perimeter := s.Perimeter()
	fmt.Println("Area: ", area)
	fmt.Println("Perimeter: ", perimeter)
	fmt.Println("r: ", r)
}