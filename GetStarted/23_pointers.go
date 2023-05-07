package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x float64
	y float64
}

// attaching a method to struct with non pointer reciever
func (v vertex) Abs() float64 {

	return math.Sqrt(v.y*v.x + v.y*v.y)
}

// rewrting abs method as function
func FuncAbs(v vertex) float64 {

	return math.Sqrt(v.y*v.x + v.y*v.y)
}

// method with pointer reciever
func (v *vertex) scale() {
	const a = 5
	v.x = v.x * a
	v.y = v.y * a

}
func (v vertex) scaleV2() {
	const a = 5
	v.x = v.x * a
	v.y = v.y * a

}

// rewriting scale() method  as a function
func FuncScale(v *vertex) {
	const a = 5
	v.x = v.x * a
	v.y = v.y * a

}
func FuncScaleV2(v vertex) {
	const a = 5
	v.x = v.x * a
	v.y = v.y * a

}

func main() {
	v1 := vertex{
		x: 5,
		y: 6,
	}
	//calling method abs: non pointer reciever
	// we call through pointer or non pointer
	r := &v1
	fmt.Println("VALUE ", v1.Abs())
	fmt.Println("VALUE 2", r.Abs())

	// calling method scale : pointer reciever : Can be called through both pointer and non pointer
	//since its pointer so values of can be modified
	//POINT1 : we can call pointer reciever method with non pointer
	//POINT2 :since its a pointer , the actually values of the struct changes
	fmt.Println("the values before are ", v1.x, v1.y)
	v1.scale()
	fmt.Println("the values after are ", v1.x, v1.y)
	fmt.Println(v1.Abs())

	//using pointer to call pointer reciver method
	p := &v1
	fmt.Println("the values BEFORE  are ", v1.x, v1.y)

	p.scale()
	fmt.Println("the values AFTER are ", v1.x, v1.y)

	//calling v2 of scaleV2 : non pointer method ,
	//it chnages value inside method , but chnages are not reflected inside struct bcoz its non pointer
	//POINT3 since its a not a pointer , the actual values of struct doesnt change

	v2 := vertex{
		x: 5,
		y: 6,
	}
	fmt.Println("the values BEFORE calling non pointer method are ", v2.x, v2.y)
	v2.scaleV2()
	fmt.Println("the values AFTER calling non pointer method are ", v2.x, v2.y) // no chnages

	//Now calling the functions instead of methods
	fmt.Println("\nNow calling the functions instead of methods")
	v1 = vertex{
		x: 5,
		y: 6,
	}
	//FuncScale(v1)
	//POINT4: cannot be called since we need to use Pointer only

	p = &v1
	FuncScale(p)

	fmt.Println("the values are ", v1.x, v1.y)

	// since this function takes only vertex (not pointer ) ,  v1 be passes  and not &v1
	fmt.Println(FuncAbs(v1))

	//calling non pointer arguments function FuncScaleV2
	fmt.Println("\ncalling non pointer arguments function FuncScaleV2")

	v1 = vertex{
		x: 5,
		y: 6,
	}
	//FuncScaleV2(v1)
	//POINT5: cannot be called since we need to use a a non pointer only
	p = &v1
	//FuncScaleV2(p)
	FuncScaleV2(v1)

	//POINT6 : Value will not chnage since it was not pointer
	fmt.Println("the values are ", v1.x, v1.y)

	/*
		Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

		while methods with pointer receivers take either a value or a pointer as the receiver when they are called:


	*/
	/*
		The equivalent thing happens in the reverse direction.

			Functions that take a value argument must take a value of that specific type:

		while methods with value receivers take either a value or a pointer as the receiver when they are called:

	*/

}
