package main

import "fmt"

func main() {

	m := map[string]int{
		"sugandha": 34,
		"manu":     38}
	fmt.Println(m)
	fmt.Println(m["manu"])
	//when the value is not found
	fmt.Println(m["sapna"]) // 0 is printed

	v, ok := m["sapna"]
	fmt.Println(v, ok) // 0 is printed
	if v, ok := m["sapna"]; ok {
		fmt.Println("the value of sapna is", m["sapna"], v) // 0 is printed
	}
	if v, ok := m["sugandha"]; ok {
		fmt.Println("the value of sugandha is", m["sugandha"], v) // 0 is printed
	}

	// create new map
	n := map[string]string{
		"sugandha": "Artist",
		"manu":     "Engineer"}
	fmt.Println(n)
	fmt.Println(n["manu"])
	fmt.Println(n["sapna"])
	//v, ok = n["sapna"] //ERROR cannot assign string to int
	_, ok = n["sapna"]

	fmt.Println(ok) // false is printed

	// adding the elemnts
	fmt.Println("Adding elements")

	n["sapna"] = "HomeChef"
	fmt.Println(n)
	// using range
	for k, v := range n {
		fmt.Println(k, " : ", v) // false is printed

	}
	// delete the elements
	//delete is a built in function delete(m,k)
	delete(n, "sapna")

	fmt.Println(n)
	delete(n, "saxena") // no error detected
	fmt.Println(n)

}
