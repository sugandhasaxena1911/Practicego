package main

import "fmt"

func main() {
	addition()
	additionV2(5, 5)
	st, num := additionV3(15, 45)
	fmt.Println(st, num)

}

// no parameters
func addition() {
	fmt.Println("hello i am addition")
}

// with parameters
func additionV2(num1 int, num2 int) {
	fmt.Println("hello i am addition with parameters")
	num3 := num1 + num2
	fmt.Println(num3)

}

// with parameters and returns
func additionV3(num1 int, num2 int) (string, int) {
	fmt.Println("hello i am addition with parameters")
	num3 := num1 + num2
	st := fmt.Sprint("the addition value is")
	return st, num3
}
