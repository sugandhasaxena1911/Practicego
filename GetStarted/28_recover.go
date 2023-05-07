package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Inside main , calling CallOpenFile")
	defer DeferInMainZ()
	CallOpenFileZ()
	fmt.Println("POINT NOTED - 3 Inside main , calling CallOpenFile -----2 ")
	defer DeferInMain2Z()

}

func CallOpenFileZ() {

	fmt.Println("hello, Calling the OPEN FILE FUNCTION ")
	defer Defer1Z()
	OpenFileZ()
	fmt.Println("POINT NOTED-2 hello, AFTER Calling OPEN FILE FUNCTION")
	defer Defer4Z()

}

func OpenFileZ() {
	fmt.Println("hello, Inside  the OPEN FILE FUNCTION ")
	defer Defer2Z()

	f, err := os.Open("ABC.txt")
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()
	defer Defer3Z()

	fmt.Println("POINT NOTED-1, Inside  the OPEN FILE FUNCTION---2 ")

}
func DeferInMainZ() {
	fmt.Println("I am defer in main ")

}
func DeferInMain2Z() {
	fmt.Println("I am defer in main---2 ")

}

func Defer1Z() {
	fmt.Println("In defer 1z")
}
func Defer2Z() {
	fmt.Println("I am defer 2 ")
	if r := recover(); r != nil {
		fmt.Println("RECOVERED NOW", r)
	}
}
func Defer3Z() {
	fmt.Println("I am defer 3 ")
}
func Defer4Z() {
	fmt.Println("I am defer 4 ")

}
