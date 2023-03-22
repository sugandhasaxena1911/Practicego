package main

import "fmt"

func main() {
	//CALLBACK is passing a func as a parameter

	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := summation_num(sl...)
	fmt.Println("Total is ", sum)

	//func to add only even number and odd numbers from list
	sum_even := summation_even(summation_num, sl...)
	fmt.Println("Total of even numbers is", sum_even)

	sum_odd := summation_odd(summation_num, sl...)
	fmt.Println("Total of odd numbers is", sum_odd)
}

func summation_num(x ...int) int {
	total := 0
	fmt.Println("I am inside addition func and i have initialized total as 0")
	for _, v := range x {

		total = total + v
	}
	return total
}

func summation_even(fn func(x ...int) int, y ...int) int {
	fmt.Println("I am inside summation_even ")
	var sl []int
	//  created a slice of even numbers
	for _, v := range y {
		if v%2 == 0 {
			sl = append(sl, v)
		}
	}
	return fn(sl...)
}

func summation_odd(fn func(x ...int) int, y ...int) int {
	fmt.Println("I am inside summation_odd ")
	var sl []int
	//  created a slice of even numbers
	for _, v := range y {
		if v%2 == 1 {
			sl = append(sl, v)
		}
	}
	return fn(sl...)
}
