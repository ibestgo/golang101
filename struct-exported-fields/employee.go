package org

// Only Employee, Firstname and Lastname are visible out of the package
type Employee struct {
	Firstname, Lastname string
	salary int
	fullTime bool
}