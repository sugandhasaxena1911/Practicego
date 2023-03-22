package main

import "fmt"

func main() {

	var x [5]int
	fmt.Println(x)
	fmt.Println(x[2])

	x[2] = 300
	fmt.Println(x)
	//length of array
	fmt.Println(len(x))
	//print the elements of array
	for y := 0; y < len(x); y++ {
		fmt.Println("The", y, " element is ", x[y])

	}
	//array of string
	var strarr [5]string
	fmt.Println(strarr)
	fmt.Println("The length od string array is ", len(strarr), " and the 2nd element is ", strarr[1])
	strarr[0] = "Hello"
	fmt.Println(strarr)

}
