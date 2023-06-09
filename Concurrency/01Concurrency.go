package main

import (
	"fmt"
	"runtime"
)

func main() {
	//another go routine spun up in micrseconds and in that time forst go routine ran

	fmt.Println("Arch ", runtime.GOARCH)
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("root ", runtime.GOROOT())
	fmt.Println("Num of go routines ", runtime.NumGoroutine())
	fmt.Println("Num of cpu ", runtime.NumCPU())

	fmt.Println("Launching another go routine ")
	go foo() ///Another go routine launched
	fmt.Println("Num of go routines ", runtime.NumGoroutine())

	bar()
}

func foo() {

	for x := 0; x < 10; x++ {
		fmt.Println("FOO X is : ", x)
	}
}

func bar() {

	for x := 0; x < 10; x++ {
		fmt.Println("BAR Y is : ", x)
	}
}
