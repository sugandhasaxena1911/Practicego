package main

import "fmt"

type intrfc interface {
	methodM()
}

type st struct {
	x, y int
}

type myfloat float64

func (s st) methodM() {
	fmt.Println("The value  x and y is ", s.x, s.y)
}

func (flt myfloat) methodM() {
	fmt.Printf("the value of my float %f %T\n", flt, flt)
}

func describe(i intrfc) {
	fmt.Printf("The value is %v %T\n", i, i)
}

func main() {
	s := st{
		x: 45,
		y: 7,
	}
	fmt.Println("Struct is ", s)
	// struct implements interface
	fmt.Println("struct implements interface")
	s.methodM()
	describe(s)

	//myfloat implements interface
	f := myfloat(7.987)
	f.methodM()
	describe(f)
}
