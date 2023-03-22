package main

import "fmt"

func main() {
	if true && true {
		fmt.Println("Hello If")
	}
	if x := 2; x == 2 {
		fmt.Println("initiation of x done in `If` to narrow down the scope of variale x")

	}
}
