package main

import(
	"fmt"
)

// our data type
type Cube struct {
	side float64
}

// interface definitions
type Shape interface {
	Area() float64
}
type Object interface {
	Volume() float64
}

// interface implementations
func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}
func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

/// misc
type Skin interface {
	Color() float64
}

func main1() {
	cube := Cube{3.0}	// OR cube := Cube{side: 3.0}

	var s Shape = cube
	var o Object = cube
	
	fmt.Println("Area  : ", s.Area())
	fmt.Println("Volume: ", o.Volume())
	fmt.Println(cube.Area())
	fmt.Println(cube.Volume())
}

func main2() {
	var s Shape = Cube{3.0}

	// This is called, Type Assertion
	// We can find out the underlying dynamic value of an interface using the syntax i.(Type)
	cube := s.(Cube)
	
	fmt.Println("Area  : ", cube.Area())
	fmt.Println("Volume: ", cube.Volume())
}

func main() {
	var s Shape = Cube{3.0}

	value1, ok1 := s.(Object)
	fmt.Printf("Value: %v, Implement interafce Object? %v\n", value1, ok1)
	fmt.Println(value1.Volume())	// OK
	//fmt.Println(value1.Area())	// Not-OK
	//fmt.Println(value1.side)		// Not-OK

	value2, ok2 := s.(Skin)
	fmt.Printf("Value: %v, Implement interafce Skin? %v\n", value2, ok2)
}