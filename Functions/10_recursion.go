package main

import "fmt"

func main() {

	fmt.Println("The factorial is", factorial(6))
	fmt.Println("The factorial is", factorialV2(6))

}
func factorial(n int) int {

	if n == 1 {
		return n
	}

	return n * factorial(n-1)
}

func factorialV2(n int) int {
	result := 1

	for x := 1; x <= n; x++ {
		result = result * x
	}

	return result

}
