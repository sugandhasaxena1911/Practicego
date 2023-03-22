package main

import (
	"fmt"
	"os"
)

func main() {

	defer test1()
	f_openfile_1()
	fmt.Println("AFTER  calling f_openfile_1 in main  ")

}

func f_openfile_1() {
	fmt.Println("BEFORE f_openfile_1  ")
	defer test2()
	f_openfile()
	fmt.Println("AFTER calling f_openfile in  f_openfile_1  ")
}

func test2() {

	fmt.Println("i am in test2 ")

}

func f_openfile() {
	n, err := os.Open("Hello.txt")
	defer f_closefile() // PS : this defer should come before panic otherwise its not called
	if err != nil {
		//fmt.Println("The error occured while opening file ", err)
		//log.Println("The error occured while opening file ", err) // gave the date and time as well, defer fnc are called
		//log.Fatalln("the error is fatal", err) //log the error and then exits  os.Exit(1), defer fnc not called
		//log.Panicln("The panic mode is on ", err) // log the error , goes to caller (call defer), and then panic()
		panic(err) //goe to caller (calld defer ) and then panic()
	}
	fmt.Println("after log ", n)
}

func test1() {

	fmt.Println("I am a test function .")
}
func f_closefile() {

	fmt.Println("i am now closed ")
}
