package main

import "fmt"

func main() {

	var x [5]int // its by default initialized with ZERO VALUE of int
	var y [5]string
	var z [5][]int
	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println(y)
	fmt.Println(y[2])
	fmt.Println("Z:  ", z) //array of empty slices
	fmt.Println(z[2])
	fmt.Printf("%T :\n", z[2])

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

	for i, v := range strarr {
		fmt.Println(i, v)
	}

}
