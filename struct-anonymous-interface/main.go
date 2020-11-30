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
	Salaried			// You can have an anonymouse interface
}

func main() {
	baechul := Employee {
		firstName: "Baechul",
		lastName: "Kim",
		
		// since anonymouse, interafce name becames a field name. 
		Salaried: Salary {1000, 500, 500},
	}
	
	// The most important thing to remember here is, 
	// in contrast with field promotion as we’ve seen earlier, 
	// only the methods are promoted when the anonymous field is an interface. 

	// U can't have baechul.basic

	fmt.Printf("Baechul's salary is %d", baechul.getSalary())
}