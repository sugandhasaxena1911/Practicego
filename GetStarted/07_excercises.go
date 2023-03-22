package main

import "fmt"

var a int = 42
var b string = "JAMES bOND"
var c bool = true

type tina int

var m tina
var n int

func main() {
	//first hands on
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(" the value are ", x, y, z)
	//second hands on
	fmt.Println(" the value are ", a, b, c)
	//use sprintf
	s := fmt.Sprintf("%v \t%v\t %v", a, b, c)
	fmt.Println(s)
	//third
	type suggu int
	var q suggu
	fmt.Println(q)
	fmt.Printf("%T\n", q)
	q = 42 //directly able to assign without conversion
	fmt.Println("value of q ", q)

	///fourth
	fmt.Println("value of m and n   ", m, n)
	fmt.Printf("type of m and n   %T %T \n", m, n)
	m = 44
	n = int(m)

	fmt.Println("value of m and n   ", m, n)
	fmt.Printf("type of m and n   %T %T \t", m, n)

}
