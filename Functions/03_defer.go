package main

import "fmt"

func main() {

	openfile1()
	defer closefile1() // it mkaes sure that file is always closed before the program exits
	otherstuff()
}
func closefile1() {
	fmt.Println("Hello , FILE1 is closed now ")
}
func openfile1() {

	fmt.Println("Hello , FILE1 is opened")
	fmt.Println("Hello , file is edited and contents are written")
	openfile2()
	defer closefile2()
}
func closefile2() {
	fmt.Println("Hello , FILE2 is close now  ")
}

func openfile2() {
	fmt.Println("Hello , FILE2 is open now  ")
	defer test()
	panic("Ohh no ..its panic time")
}

func test() {
	fmt.Println("TEST LAST DEFER WHERE PANIC OCCURS ")
}

func otherstuff() {
	fmt.Println("continue with other processing ")
}
