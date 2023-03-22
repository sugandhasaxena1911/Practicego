package main

import "fmt"

func main() {
	fn1 := incrementor()
	fn2 := incrementor()
	fmt.Println(fn1())
	fmt.Println(fn1())
	fmt.Println(fn1())
	fmt.Println(fn2())
	fmt.Println(fn2())
	fmt.Println(fn2())
	fn2 = incrementor()
	fmt.Println(fn2())

}
func incrementor() func() int {
	fmt.Println("Hi , i am inside incrementor function")
	var x int
	return func() int {
		fmt.Println("Hi , i am inside returned function")
		x++
		return x
	}

}
