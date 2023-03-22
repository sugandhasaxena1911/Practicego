package main

import "fmt"

func main() {
	//   var <varname> <vartype>, type <typename> <underlying type>
	var a int = 5
	fmt.Println("Value of a ", a)
	type suggu int
	var b suggu
	//cannot use a (variable of type int) as type suggu in assignment
	//b = a

	//CONVERSION
	b = suggu(a)
	b = 100 // directly able to assign without conversion
	//a = b   erorr
	a = int(b)
	fmt.Println("Value of b ", b)

}
