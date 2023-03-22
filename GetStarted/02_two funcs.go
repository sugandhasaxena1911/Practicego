package main

import "fmt"

func main() {
	fmt.Println("Hello, Im in main ")
	foo()
}

func foo() {
	fmt.Println("Im in foo ")
}
