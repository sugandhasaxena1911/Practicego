package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("hello")
	//https://go.dev/play/p/HVc6Wp9MiEN
	var bs = []byte(`[{"Name":"vedanshi","Age":2.5,"Toys":["puzzle","painting","balloons"]},{"Name":"Vedansh","Age":1,"Toys":["puzzle","painting","maths"]}]
`)
	fmt.Println("Hello ji ", string(bs))
	type baby struct {
		Name string
		Age  float64
		Toys []string
	}
	var babies []baby
	//you an also use below
	//babies := []baby{}
	err := json.Unmarshal(bs, &babies)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range babies {
		fmt.Println("Index ", i)
		fmt.Println("Value ", v)
		fmt.Println("Value is ", babies[i])
		fmt.Println("value is ", v.Name, v.Age, v.Toys)
	}
}
