/*
Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types.
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int		`json:"page"`
	Fruits []string	`json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)	// Marshal returns the JSON encoding of v.
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))		// "gopher"
	
	slcD := []string {"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))		// ["apple","peach","pear"]
	
	mapD := map[string]int {"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))		// {"apple":5,"lettuce":7}
	
	//////
	res1D := &response1 {
        Page:   1,
        Fruits: []string {"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))		// {"Page":1,"Fruits":["apple","peach","pear"]}

    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))		// {"page":1,"fruits":["apple","peach","pear"]}
	
	///// decoding to map
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
	fmt.Println(dat)				// map[num:6.13 strs:[a b]]
	
	// In order to use the values in the decoded map, we’ll need to convert them to their appropriate type.
	num := dat["num"].(float64)		
	fmt.Println(num)
	
	strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
	fmt.Println(str1)
	
	//////
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)				// {1 [apple peach]}
	fmt.Println(res.Fruits[0])		// apple
	
	//+++
	abc := response2 {
		Page: 3,
		Fruits: []string {"pear", "blueberry"},
	}
	fmt.Println(abc)				// {3 [pear blueberry]}
	//+++

	/*
	In the examples above we always used bytes and strings as intermediates between the data and JSON representation 
	on standard out. We can also stream JSON encodings directly to os.Writers like os.Stdout 
	or even HTTP response bodies.
	*/
	enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)					// {"apple":5,"lettuce":7}
}