package main

import (
	"fmt"
)

type Contact struct {
	phone, address string
}
type Employee struct {
	name string
	salary int
	Contact		// anonymous field. Same as Contact: Contact
}

// method on the parent struct
func (e *Employee) changePhone(newPhone string) {
	e.phone = newPhone
	//e.Contact.phone = newPhone	// OR this
}

func main() {
	e := Employee {
		name: "Baechul",
		salary: 12000,
		Contact: Contact {"000-000-0000", "Abc street, USA"},
	}

	fmt.Println("Before\t", e.phone) // phone is promoted to parent when anonymous field is used
	e.changePhone("408-123-4567")
	fmt.Println("After\t", e.Contact.phone) // OR e.phone
}