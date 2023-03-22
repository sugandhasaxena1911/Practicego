package main

import "fmt"

type employee struct {
	firstname string
	lastname  string
	age       int
}

// create method
//Go language support methods.
//Go methods are similar to Go function with one difference, i.e, the method contains a receiver argument in it.
//With the help of the receiver argument, the method can access the properties of the receiver
//this is method with struct type reciever

func (emp employee) profession(str string) {
	fmt.Println("My name is", emp.firstname, emp.lastname, "and profession is ", str)
}

func main() {

	emp1 := employee{
		firstname: "sugandha",
		lastname:  "saxena",
		age:       34,
	}

	emp1.profession("Artist")
	//unable to call directly because the method is binded with the reciever
	//profession("Doctor")

}
