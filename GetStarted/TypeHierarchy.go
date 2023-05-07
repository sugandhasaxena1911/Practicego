package main

import "fmt"

type st1 struct {
	n int
	m int
}

// embedded struct
type st2 struct {
	st1
	p int
}

func main() {

	var a st1
	b := st2{
		st1{7, 8},
		9,
	}
	fmt.Println(a, b)
	// superclass variable is assigned the value of subclass : Possible in java , But its not acceptable for Go
	// Go Doesnot support type hierarchy
	//a = b

	//posible
	a = b.st1

	//However subclass may call methods of superclass
	b.f1()

	//callig f2 , kind of overriden function . but since no type hierarchy thefore no runtime polymorph like this
	// for runtime poly morph use : Interfacses
	a.f2()
	b.f2()

}

func (s st1) f1() {
	fmt.Println("im method of st1-F1", s.n, s.m)
}

// f2 defined for both
func (s st1) f2() {
	fmt.Println("im method of st1-F2", s.n, s.m)
}
func (s st2) f2() {
	fmt.Println("im method of st2-F2", s.n, s.m, s.p)
}
