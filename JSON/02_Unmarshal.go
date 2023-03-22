package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var j = []byte(`[{"Name":"vedanshi","Age":2.5,"Toys":["puzzle","painting","balloons"]},
                     {"Name":"Vedansh","Age":1,"Toys":["puzzle","painting","maths"]}
					]`)
	type baby struct {
		Name string
		Age  float64
		Toys []string
	}
	var baby1 []baby
	err := json.Unmarshal(j, &baby1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(baby1)
	fmt.Printf("%v\n", baby1)
	fmt.Printf("%+v", baby1) //when printing structs, the plus flag (%+v) adds field names

}
