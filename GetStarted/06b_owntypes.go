package main

import "fmt"

func main() {
	var a int
	type hotdog int
	var b hotdog
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

}
