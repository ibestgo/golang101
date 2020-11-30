package main

import (
	"fmt"
	"strings"
)

// MyString and MyInteger are custom data types
type MyString string
type MyInteger int
type Rect struct {
	width, height float64
}

// Note tha it accepts an empty interface. All types implement empty interface implicitly.
// hence we can pass any data type to this method.
func describe(i interface{}) {
	fmt.Printf("Data type is %T and valus is %v\n", i, i)
}

func describe2(i interface{}) {
	switch i.(type) {	// Type Switch pattern. <interface>.(type) only for switch staements
	case string:
		fmt.Println("i stored string", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i stored int", i)
	default:
		fmt.Println("i stored something else", i)			
	}
}

func main2() {
	describe2("Hello World")
	describe2(52)
	describe2(true)
}

func main1() {
	ms := MyString("Hello World")
	mi := MyInteger(5)
	r := Rect{5.5, 4.5}
	describe(ms)
	describe(mi)
	describe(r)
}

func main() {
	main1()
	main2()
}