package main

import "fmt"

func main() {
	//simple case
	switch {
	case true:
		fmt.Println("this is true")
	case false:
		fmt.Println("This is false")
	}

	// default case
	switch "Hello" {
	case "Jello":
		fmt.Println("this is true too")
	case "Gello":
		fmt.Println("This is false")
	default:
		fmt.Println("prints default")

	}

	//fallthrough : not recommended
	switch {
	case true:
		fmt.Println("this is true 3")
		fallthrough
	case 2 == 2:
		fmt.Println("This is true 4 ")
		fallthrough
	default:
		fmt.Println("prints default")

	}

	// switch case
	c := 5

	switch c {
	case 1, 2, 3:
		fmt.Println("This is 1 2 3 ")
	case 4, 5, 6:
		fmt.Println("This is 4 5 6 ")
	default:
		fmt.Println("This is 7 8 9 10")

	}

}
