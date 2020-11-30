package main

// types is new in v2.0.1 release
// . updated go.mod to the newer version (require github.com/baechul/golang/v2 v2.0.1)
// . then run "go get -u github.com/baechul/golang/v2" to download the latest v2.0.1
import (
	"fmt"
	"github.com/baechul/golang/v2/types"
)

// Request shows anonymous and nested struct example.
type Request struct {
	types.NamespacedName	// anonymous field. the key will be "NamespacedName"
}

func main() {
	req := Request {
		NamespacedName: types.NamespacedName {	// ***
			Namespace: "default",
			Name: "haha",
		},
	}

	fmt.Println(req)				// {{default haha}}
	fmt.Println(req.NamespacedName)	// {default haha}
}