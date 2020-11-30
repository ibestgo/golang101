package main

import (
	"fmt"
	"org"
)


func main() {
	baechul := Employee {
		firstName: "Baechul",
		lastName: "Kim",
		
		// since anonymouse, interafce name becames a field name. 
		Salaried: Salary {1000, 500, 500},
	}

	fmt.Println(baechul)
}