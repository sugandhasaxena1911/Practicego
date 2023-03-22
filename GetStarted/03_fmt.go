package main

import "fmt"

func main() {
	m, e := fmt.Println("arguement 1 ", "arg2", true, 1234)
	n, _ := fmt.Println("arg1", "arg2", true, 1234)
	fmt.Println(m, e, n)
	//Please see that you have to use n , e variables otherwise code will not compile
	// You can use underscore : _     if you dont want to use in anywhere in the code

	//short declaration operator
	abc := "hello"
	fmt.Println("print me ", abc)
	abc = "chnaged hello" // gives error if short declration operator is used
	// no new variables on left side of :=
	fmt.Println("abc ", abc)

	// checking the printf , -->it doesnot append the /n at the end   ,
	//its known as format printing
	fmt.Printf("Hello, im printf ")
	fmt.Printf("Hello, im printf , 2")
	fmt.Printf("Hello, im printf , 3\n")
	// verbs used for print  : %v, %#v , %T  ,%%   , %t ,   %b , %d  , %o, %x  , %X
	txt := "Im am text"
	fmt.Printf("%v \n", txt)
	fmt.Printf("%#v \n", txt)
	fmt.Printf("%T \n", txt)
	//fmt.Printf("%% \n", txt)
	nm := 911
	fmt.Printf("%b \n", nm) //binary
	fmt.Printf("%d \n", nm) //decimal
	fmt.Printf("%o \n", nm) //octal
	fmt.Printf("%x \n", nm) //hexadecimal
	fmt.Printf("%X \n", nm)
	bool := false
	fmt.Printf("%t \t %v \n", bool, bool) //binary

	//sprint , sprintf , sprintln
	txt2 := fmt.Sprintln(txt, "Hello")
	txt2 = fmt.Sprintf("%v%v", txt, "\tHello")

	fmt.Println(txt2)

}
