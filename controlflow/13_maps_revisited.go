package main

import "fmt"

func main() {

	var m1 map[string]int
	m2 := make(map[string]int)
	m3 := make(map[string]int, 2)

	//m1["tina"] = 1 // gives error : assignment to entry in nil map
	m2["tina"] = 1
	m3["tina"] = 1
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	//The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them
	m3["sugandha"] = 2
	m3["manu"] = 3
	fmt.Println(len(m3))

	var m4 = make(map[string]int, 5) // note the equality sign and not :=  , because  var is used
	m4["tina"] = 1
	fmt.Println(m4)

}
