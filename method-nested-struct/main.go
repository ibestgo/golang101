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
	contact Contact
}

// method on the parent struct
func (e *Employee) changePhone(newPhone string) {
	e.contact.phone = newPhone
}

// Or you can have a method on the nested struct
func (c *Contact) changePhone(newPhone string) {
	c.phone = newPhone
}

func main2() {
	e := Employee {
		name: "Baechul",
		salary: 12000,
		contact: Contact {"000-000-0000", "Abc street, USA"},
	}

	fmt.Println("Before\t", e.contact.phone)
	e.contact.changePhone("408-123-4567")
	fmt.Println("After\t", e.contact.phone)
}

func main1() {
	e := Employee {
		name: "Baechul",
		salary: 12000,
		contact: Contact {"000-000-0000", "Abc street, USA"},
	}

	fmt.Println("Before\t", e.contact.phone)
	e.changePhone("408-123-4567")
	fmt.Println("After\t", e.contact.phone)
}

func main() {
	main1()
	main2()
}

