package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer test()
	n, err := os.Open("Hello.txt")
	if err != nil {
		//fmt.Println("The error occured while opening file ", err)
		//log.Println("The error occured while opening file ", err) // gave the date and time as well, defer fnc are called
		//log.Fatalln("the error is fatal", err) //log the error and then exits  os.Exit(1), defer fnc not called
		log.Panicln("The panic mode is on ", err)
	}
	fmt.Println("after log ", n)
}

func test() {

	fmt.Println("I am a test function .")
}
