package main

import "fmt"

func main() {

	sum := summation(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)
	//unfurling the slice
	x := []int{10, 11, 12, 12, 14, 15}
	//sum =  summation(x)  cannot use x directly
	sum = summation(x...)
	fmt.Println(sum)

	//  what if you dont pass parameter bcoz for variadic means 0 or more
	// slice is created , but its empty ie underlying array is not created
	sum = summation()
	fmt.Println(sum)
}

func summation(x ...int) int {
	fmt.Println(x)        // printed as slice
	fmt.Printf("%T\n", x) //slice of int

	sum := 0

	for _, v := range x {
		sum = sum + v
	}

	return sum

}
