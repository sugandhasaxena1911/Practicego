package main

import "fmt"

func main() {

	const a = 42
	const b = "hello"
	const (
		c = 67
		d = "Hello ji "
	)

	fmt.Printf("value of a is %v and type is  %T \n", a, a)
	fmt.Printf("value of b is %v and type is  %T \n", b, b)
	fmt.Printf("value of c is %v and type is  %T \n", c, c)
	fmt.Printf("value of d is %v and type is  %T", d, d)

}
