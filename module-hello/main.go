package main

import (
	"fmt"
	"github.com/baechul/golang/calc"
	// https://github.com/baechul/golang/calc returns 404 Not Found. 
	// How Go resolves the package import URL? My guess is that since GitHub is 
	// one of the recognized code hosting sites, it knows where the package is located

	// v2 is a new branch and treat like another dependency.
	calcNew "github.com/baechul/golang/v2/calc"
)

func main() {
	result := calc.Add(3, 4, 3)
	fmt.Println("calc.Add(3, 4, 3) => ", result)

	// v2
	if newResult, err := calcNew.Add(3, 4, 3); err != nil {
		fmt.Println("Error: => ", err)
	} else {
		fmt.Println("calcNew.Add(3, 4, 3) => ", newResult)
	}

	// newResult, err := calcNew.Add(3, 4, 3)
	// if err != nil {
	// 	fmt.Println("Error: => ", err)
	// } else {
	// 	fmt.Println("calcNew.Add(3, 4, 3) => ", newResult)
	// }

}