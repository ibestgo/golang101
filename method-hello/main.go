package main

import (
	"fmt"
)

type Employee struct {
	Firstname, Lastname string
}

// this is a function NOT a method yet
func fullname(fname string, lname string) (fullname string) {
	fullname = lname + ", "+fname	// note that we used fullname func parameter for return 
	return	// No need of "return fullname" as we used fullname func parameter.
}

func (e Employee) fullname() string {
	return e.Lastname + ", " + e.Firstname
}

func main() {
	e := Employee {
		Firstname: "Baechul",
		Lastname: "Kim",
	}

	fmt.Println(fullname(e.Firstname, e.Lastname))
	// the problem: for 10000 employee, first/last names have to be passed.
	// how about just e.fullname()?

	fmt.Println(e.fullname())
}

