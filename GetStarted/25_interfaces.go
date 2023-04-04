package main

import "fmt"

type intrfc interface {
	methodM()
}

type st struct {
	x, y int
}

type myfloat float64

func (s *st) methodM() {
	if s == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println("The value  x and y is ", s.x, s.y)
}

func (flt myfloat) methodM() {
	fmt.Printf("the value of my float %f %T\n", flt, flt)
}

func describe(i intrfc) {
	fmt.Printf("inside describe The value is %v %T\n", i, i)
}

func main() {
	s := st{
		x: 45,
		y: 7,
	}
	fmt.Println("Struct is ", s)
	// struct implements interface
	fmt.Println("struct implements interface")
	//methodM() can be called  by both st and &st, but we cannot describe() where argument value is interface
	s.methodM()
	//describe(s)
	describe(&s)

	//myfloat implements interface
	fmt.Println("myfloat implements interface")

	f := myfloat(7.987)
	f.methodM()
	describe(f)

	/// interfaces value with nil underlying value
	fmt.Println(" interfaces value with nil underlying value")
	s = st{} //non nil
	var b st //non nil
	var c *st
	fmt.Printf("the vlaue and type %v %T\n", s, s)
	fmt.Printf("the vlaue and type %v %T\n", b, b)
	fmt.Printf("the vlaue and type %v %T\n", c, c)

	fmt.Println(" calling with interface value whcic has non nil nalue ")
	s.methodM()
	describe(&s)

	b.methodM()
	describe(&b)

	fmt.Println(" calling with interface value which has  nil nalue ")
	c.methodM() // throw error since c is nil and it cnnot print nil elemnets value
	//hefore ad one handling line
	//describe(s)
	describe(c)

	///calling with nil interface value
	fmt.Println("\t///calling with nil interface value\n")
	var i intrfc
	describe(i)

	//Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	i.methodM()
}
