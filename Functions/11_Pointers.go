package main

import "fmt"

func main() {

	a := 43
	fmt.Println("the value of a : ", a)
	fmt.Println("the address of a : ", &a)
	fmt.Printf("the type of the address %T \n: ", &a) // *int

	fmt.Println("the value  of a : ", *&a)
	*&a = 55
	fmt.Println("the value  of a : ", a)
	b := &a
	*b = 54
	fmt.Println("the value  of a : ", a)

	st := "Hello"
	fmt.Printf("the type of the address %T \n: ", &st)

	//using pointers as parameter
	fmt.Println("checking pointer as parameter ")
	fmt.Println("The value of st before ", st)
	passbyreference(&st)
	fmt.Println("The value of st after ", st)

}

func passbyreference(a *string) {
	fmt.Println("the passed parameter", a)
	*a = "hello sugandha"

}
