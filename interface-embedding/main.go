package main

import (
	"fmt"
)

type Cube struct {
	side float64
}

type Shape interface {
	Area() float64
}
type Object interface {
	Volume() float64
}

// an interface cannot implement other interfaces or extend them, 
// but we can create a new interface by merging two or more.
type Material interface {
	Shape
	Object
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}
func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	c := Cube {3}
	var m Material = c
	var s Shape = c
	var o Object = c
	fmt.Printf("Dynamic type and value of static type Material: %T, %v\n", m, m)
	fmt.Printf("Dynamic type and value of static type Shape: %T, %v\n", s, s)
	fmt.Printf("Dynamic type and value of static type Object: %T, %v\n", o, o)
}