package main

import "fmt"

func main() {

	//https://www.programiz.com/golang/anonymous-function

	profession("Artist")

	//Anonymous functions hav no identifiers
	func() {
		fmt.Println("I am anonymous")
	}()
	func(str string) {
		fmt.Println("I am anonymous-2 ", str)
	}("Coding")

	// we can assign the anonymous func to variable
	area := func(l int, b int) {
		fmt.Println("the area of the rectangle is ", l*b)
	}
	area(34, 2) //function call

	//Anonymous functions returning values
	volume := func(l, b, h int) int {
		return l * b * h
	}

	v := volume(2, 5, 12)
	fmt.Println("the volume of cuboid is ", v)

	// you can use directly as well
	fmt.Println("the volume of cuboid is ", volume(5, 7, 8))

	//Anonymous functions as an argument to another function    :
	getheight := func(h float64) float64 {
		return h * 2
	}

	fmt.Println("the volume of cylinder is ", cylinder_vol(getheight(3), 5.5))

	// regular func that returns anonymous func
	x := regularfunc()
	x()

	fmt.Println("Calling v2")

	y := regularfuncv2()
	fmt.Printf("%T\n", y)
	fmt.Println("the value is ", y())
	fmt.Println("the value is ", regularfuncv2()())

}

func profession(str string) {

	fmt.Println("My profession is ", str)
}

func cylinder_vol(h float64, r float64) float64 {
	var pi float64
	var vol float64
	pi = 3.14
	vol = pi * r * r * h
	return vol
}

// anonymous func as a return
func regularfunc() func() {
	fmt.Println("i am regular func but returns an anoymous func ")
	return func() {
		fmt.Println("i am anonymous as return")
	}
}
func regularfuncv2() func() int {
	fmt.Println("i am regular func but returns an anonymous func ")
	return func() int {
		fmt.Println("i am anonymous function as return , which in turn returns int")
		return 256
	}
}
