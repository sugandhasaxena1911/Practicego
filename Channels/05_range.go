package main

import "fmt"

func main() {
	ch := make(chan int)
	go fibonacciSender(10, ch) // NOTE : this is generic to specific
	//RECIEVER
	for x := range ch {
		fmt.Println(x)
	}

}

func fibonacciSender(n int, ch chan<- int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {

		ch <- x
		x, y = y, x+y
	}
	// comment and then check output : only after printing everything it gives error
	//Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
	close(ch)
}
