package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// we have introduces waitgroups now

	fmt.Println("Arch ", runtime.GOARCH)
	fmt.Println("OS ", runtime.GOOS)
	fmt.Println("root ", runtime.GOROOT())
	fmt.Println("Num of go routines ", runtime.NumGoroutine())
	fmt.Println("Num of cpu ", runtime.NumCPU())

	fmt.Println("\n\nLaunching another go routine FOO ")
	wg.Add(1)
	go foo1() ///Another go routine launched
	fmt.Println("Num of go routines FOO ", runtime.NumGoroutine())
	fmt.Println("Num of cpu FOO ", runtime.NumCPU())

	bar1()
	fmt.Println("\n\nNum of go routines BAR ", runtime.NumGoroutine())
	fmt.Println("Num of cpu BAR ", runtime.NumCPU())

	wg.Wait()
}

func foo1() {

	for x := 0; x < 10; x++ {
		fmt.Println("FOO is : ", x)
	}
	wg.Done()
	fmt.Println("Nothing left to wait now  ")
}

func bar1() {

	for x := 0; x < 10; x++ {
		fmt.Println("BAR is : ", x)
	}
}
