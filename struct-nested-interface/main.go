package main

import (
	"fmt"
)

// Salary struct has only properties and no methods.
type Salary struct {
	basic 		int
	insurance	int
	allowance	int
}

// So now I want Salary struct has a method to calculate the salary. Then use an interface.

// Salaried interface defines the methods.
type Salaried interface {
	getSalary() int
}

///// How to make Salary implement the methods of the interface?
func (s Salary) getSalary() int {
	return s.basic + s.insurance + s.allowance
}

// Now let's create Employee struct that has a salary.

// Employee ...
type Employee struct {
	firstName, lastName string
	salary Salaried		// note that instead of int, we let it calculated via an interface
	// salary Salary	// you could do a specific struct but this is not for polymorphism.
}

func main() {
	baechul := Employee {
		firstName: "Baechul",
		lastName: "Kim",
		salary: Salary {1000, 500, 500},	// , tells go compiler it is the end of the stmt hence required.
	}

	fmt.Printf("Baechul's salary is %d", baechul.salary.getSalary())
	// retStr := fmt.Sprintf("Baechul's salary is %d", baechul.salary.getSalary())
	// fmt.Printf(retStr)
}