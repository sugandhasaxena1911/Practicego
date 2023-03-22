package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// create struct

	type baby struct {
		Name string
		Age  float64
		Toys []string
	}
	baby1 := baby{
		Name: "vedanshi",
		Age:  2.5,
		Toys: []string{"puzzle", "painting", "balloons"},
	}
	fmt.Println(baby1)
	// marshal
	b, err := json.Marshal(baby1)
	if err != nil {
		log.Fatalln("opps ", err)
	}
	fmt.Printf("%T\n", b)
	os.Stdout.Write(b)
	fmt.Println(string(b))
	fmt.Println(b)

	//what if i want 2 babies in json ?
	fmt.Println("Two babies data  in json")
	baby2 := baby{
		Name: "Vedansh",
		Age:  1,
		Toys: []string{"puzzle", "painting", "maths"},
	}
	babies := []baby{baby1, baby2}
	bs, err2 := json.Marshal(babies)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(bs))

}
