package main

import "fmt"

func main() {
	var a int
	var b float64

	x := 42   //int
	y := 42.5 //float64
	//explicit type declartion done above
	a = 45
	b = 45.67

	fmt.Printf("x value and type %v \t %T\n", x, x)
	fmt.Printf("y value and type %v \t %T \n ", y, y)
	fmt.Printf("a value and type %v \t %T\n", a, a)
	fmt.Printf("b value and type %v \t %T\n  ", b, b)
	// int8 : set of all signed 8-bit integers   -128 to 127
	//uint8 :  unsigned 8 bit integers : 0 to 255
	//similarly we have int16, int32, int64  & uint16, uint32, uint64

	var z int8
	z = 125 //this is not in range
	fmt.Printf("z value and type %v \t %T\n", z, z)

	//float32 and float 64  : iEEE-754 32 bit & 64 floating point numbers
	//byte and rune : alias for uint8 and int32

	// when you use only int , uint : depend upon the cpompiler presision 32 or 64 bit, int same as uint
	var c int = 9223372036854775807   //increase by 1 and you will get error
	var d uint = 18446744073709551615 //increase by 1 and you will get error
	fmt.Printf("int c value and type %v \t %T\n", c, c)
	fmt.Printf("unit d value and type %v \t %T\n", d, d)

}
