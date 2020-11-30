package main

import (
	"fmt"
)

type Employee struct {
	name string
	salary int
}

// method on the parent struct
func (e *Employee) changeName(newName string) {
	e.name = newName
}

func (e Employee) showSalary() {
	e.salary = 1500
	fmt.Println("Salary = ", e.salary)
}

func main() {
	e := Employee {
		name: "Baechul",
		salary: 12000,
	}

	fmt.Println("Before\t", e)

	e.changeName("Debora")
	e.showSalary()
	(&e).showSalary()	// &e is an address i.e a pointer

	fmt.Println("After\t", e)
}