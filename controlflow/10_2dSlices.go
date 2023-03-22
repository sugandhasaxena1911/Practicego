package main

import "fmt"

func main() {

	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("The values")
	for rows := 0; rows < len(a); rows++ {
		for cols := 0; cols < len(a[rows]); cols++ {
			fmt.Println(a[rows][cols])

		}
	}

	b := []int{10, 11, 12}
	c := []int{13, 14, 15, 16}
	d := [][]int{b, c}
	fmt.Println(d)

}
