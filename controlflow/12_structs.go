package main

import "fmt"

func main() {
	//aggregate data type , complex data type, composite data type
	type person struct {
		first_name string
		last_name  string
		age        int
	}
	// not good programmimg style , where fields assignment is not clear
	person1 := person{
		"sugandha",
		"saxena",
		34}
	fmt.Println(person1)

	//another way to assign values: clear coding
	person2 := person{
		first_name: "manu",
		last_name:  "saxena",
		age:        38}
	fmt.Println(person2)
	//access individual fields
	fmt.Println(person1.first_name, person1.last_name)
	fmt.Println(person2.first_name, person2.last_name)

	// embedded structs
	type content_creator struct {
		person
		expertise string
	}
	cc1 := content_creator{
		person: person{
			first_name: "sugandha",
			last_name:  "saxena",
			age:        34,
		},
		expertise: "artist",
	}
	fmt.Println("content creator", cc1)

	cc2 := content_creator{
		person:    person2,
		expertise: "engineer",
	}
	fmt.Println("content creator", cc2)

	//access individual fields  : directly : because fields get promoted to uppertype from intertype
	fmt.Println(cc1.first_name, cc1.last_name, cc1.age, cc1.expertise)

	//access individual fields  : indirectly  : though its permitted to use like this as well
	fmt.Println(cc2.person.first_name, cc2.person.last_name, cc2.person.age, cc2.expertise)

	//Anonymous structs
	//where you just want strucy usage on a very small area & you dont need it to use it often

	person3 := struct {
		name string
		age  int
	}{
		"sugandha",
		34,
	}
	fmt.Println(person3)
}
