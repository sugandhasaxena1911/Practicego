package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Inside main , calling CallOpenFile")
	defer DeferInMain()
	CallOpenFile()
	fmt.Println("POINT NOTED - 3 Inside main , calling CallOpenFile -----2 ")
	defer DeferInMain2()

}

func DeferInMain() {
	fmt.Println("I am defer in main ")

}
func DeferInMain2() {
	fmt.Println("I am defer in main---2 ")

}

func Defer1() {
	fmt.Println("I am defer 1 ")
}
func Defer2() {
	fmt.Println("I am defer 2 ")
}
func Defer3() {
	fmt.Println("I am defer 3 ")
}
func Defer4() {
	fmt.Println("I am defer 4 ")
}

func CallOpenFile() {

	fmt.Println("hello, Calling the OPEN FILE FUNCTION ")
	defer Defer1()
	OpenFile()
	fmt.Println("POINT NOTED-2 hello, AFTER Calling OPEN FILE FUNCTION")
	defer Defer4()

}

func OpenFile() {
	fmt.Println("hello, Inside  the OPEN FILE FUNCTION ")
	defer Defer2()

	f, err := os.Open("ABC.txt")
	if err != nil {
		//log.Println(err) //prints the error ,
		//All POINT NOTED is printed
		//all defers are called

		//log.Panicln(err)
		// Println called : error printed
		// then calls panic()
		//interrupts normal execution , defer 3 not called, POINT NOTED, is not printed.
		// calls the defer functions that were deferred before , defer 2 called
		// goes to caller , behaves as if panic happend
		// defer 4 not called, POINT NOTED-2, is not printed.
		// goes to caller main : POINT NOTED - 3  is not printed , DeferInMain() called , DeferInMain2 not called
		// then prints panic and error , prints the error route followed ,  and exits with exit code-2

		log.Fatalln(err)
		//prints the error
		//exits with exit code-1
		//Print noted is not called --ALL
		//Functions defer  are not called
	}
	defer f.Close()
	defer Defer3()

	fmt.Println("POINT NOTED-1, Inside  the OPEN FILE FUNCTION---2 ")

}
