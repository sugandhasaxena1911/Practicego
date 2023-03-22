package main

import "fmt"

func main() {
	var t string = "Hello string /n string2"
	fmt.Println("Hello ", t, "%T", t)
	// go is static typed
	//t = 123    error
	//raw literal, its taking the string as is with /n etc .
	var y = `Hello my name is `
	fmt.Println(" value Y", y)
	y = `Hello my name is 
sugandha   !! @ 
	saxena`
	fmt.Println(" value Y", y)
}
