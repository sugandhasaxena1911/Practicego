package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println("value of i ", i)
	}

	//nested for loop

	for j := 1; j <= 5; j++ {
		fmt.Println("the outer loop ", j)
		for k := 1; k <= 3; k++ {
			fmt.Println("\t the inner loop ", k)

		}
	}
}
