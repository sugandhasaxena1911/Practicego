package main

import "fmt"

var z = 30 // using var keyword outside main body
var q int  // var key word outside, but zero value used

func main() {
	x := 40
	var y = 50
	fmt.Println("Hello X", x)
	fmt.Println("Hello Y", y)
	fmt.Println("Hello Z", z)
	fmt.Println("Hello Q", q)
	foo2()

}

// foo redeclared in this package
func foo2() {
	fmt.Println("Im in foo")
	// fmt.Println("Hello X", x)  // x , y has no scope in this func
	// fmt.Println("Hello Y", y)
	fmt.Println("Hello Z", z)
	fmt.Println("Hello Q", q)
}
