package main

import "fmt"

func main() {
	send_ch := make(chan<- int, 1) //send only channel
	fmt.Printf("%T\n", send_ch)
	send_ch <- 42
	//fmt.Println("Reciev from channel ", <-send_ch) //invalid operation: cannot receive from send-only channel send_ch

	// general to specific
	sch := make(chan<- int)
	rch := make(<-chan int)
	ch := make(chan int)
	fmt.Printf("%T %T %T\n", sch, rch, ch) //chan<- int <-chan int chan int
	sch = ch                               //general to specific allowed
	//ch = rch // not allowed specific to general
}
