package main

import "fmt"

type intrfc1 interface {
	methodM()
}

type st2 struct {
	x, y int
}

type myfloatA float64

func (s *st2) methodM() {
	if s == nil {
		fmt.Println("Printting <nil> that saves us from exception")
		return
	}

	fmt.Println("The value  x and y is ", s.x, s.y)
}

func (flt myfloatA) methodM() {
	fmt.Printf("the value of my float %f %T\n", flt, flt)
}

func describe(i intrfc1) {
	fmt.Printf("inside describe The value is %v %T\n", i, i)
	fmt.Println("Calling interface methods ")
	i.methodM()
}

func main() {
	s := st2{
		x: 45,
		y: 7,
	}
	fmt.Println("Struct is ", s)
	// struct pointer implements interface
	fmt.Println("struct pointer implements interface")
	//methodM() can be called  by both st2 and &st2,its an ease that golang provides to us
	////but we cannot pass : s inside describe() or we can assign only &st to interface variable
	s.methodM()
	q := &s
	q.methodM()
	//describe(s)
	describe(&s)

	//myfloatA implements interface
	fmt.Println("myfloatA implements interface")

	f := myfloatA(7.987)
	f.methodM()
	describe(f)
	fmt.Println("using pointer for non pointer reviver method ")
	(&f).methodM()
	fmt.Println("Passing pointer t describe : interface variable is of type &f") // WHY ??????????
	r := &f
	describe(r)

	/// interfaces value with nil underlying value
	fmt.Println(" interfaces value with nil underlying value")
	s = st2{}   //non nil
	point := &s //non nil
	var b st2   //non nil
	var c *st2  //nil
	fmt.Printf("the vlaue and type %v %T\n", s, s)
	fmt.Printf("the vlaue and type %v %T\n", b, b)
	fmt.Printf("the vlaue and type %v %T\n", c, c)
	fmt.Printf("the vlaue and type %v %T\n", point, point)

	fmt.Println(" calling with interface value whcic has non nil nalue ")
	s.methodM()
	describe(&s)

	b.methodM()
	describe(&b)

	fmt.Println(" calling with non nill pointer ")
	point.methodM()
	describe(point)
	fmt.Println("Calling with nill pointer")
	c.methodM() // throw error since c is nil and it cnnot print nil elemnets value
	//thefore add below  one handling line inside methodM()
	/*if s == nil {
		fmt.Println("<nil>")
		return
	}*/
	describe(c)

	///calling with nil interface value
	fmt.Println("\t///calling with nil interface value\n")
	var i intrfc1 //nil interface
	describe(i)   //panic: runtime error: invalid memory address or nil pointer dereference

	//Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
	//i.methodM() //panic: runtime error: invalid memory address or nil pointer dereference

}
