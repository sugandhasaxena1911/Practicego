package main

import "fmt"

// create struct Person
type person struct {
	firstname string
	lastname  string
}

type employee struct {
	person
	empId      int
	department string
}

type myinterface interface {
	describeMe()
}

func main() {
	person1 := person{
		firstname: "Sugandha",
		lastname:  "Saxena",
	}
	fmt.Println(person1)
	//calling its own METHOD
	person1.speaks()
	//Calling embedded struct
	emp1 := employee{
		person:     person1,
		empId:      123,
		department: "Developer",
	}
	fmt.Println(emp1)
	emp1.speaks()

	// emp1, person1 is of type myinterface
	emp1.describeMe()
	person1.describeMe()

	//calling human, polymorphism
	human(emp1)
	human(person1)
}

// defining a method for person
func (p person) speaks() {
	fmt.Println("hello, my name is ", p.firstname, p.lastname)
}

/*func (e employee) speaks() {
	fmt.Println("hello , my name is ", e.firstname, e.lastname, " and department is : ", e.department)
}*/

func (p person) describeMe() {
	fmt.Println("this is my description . I AM PERSON")
}
func (e employee) describeMe() {
	fmt.Println("this is my description . I AM EMPLOYEE")
}

func human(intfc myinterface) {
	fmt.Println("I AM HUMAN", intfc)
	switch intfc.(type) {
	case person:
		fmt.Println("Hello PERSON", intfc.(person).firstname)
	case employee:
		fmt.Println("Hello EMPLYOEE", intfc.(employee).department, intfc.(employee).firstname)
	}

}
