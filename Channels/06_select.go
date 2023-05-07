package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go sender(odd, even, quit)
	reciever(odd, even, quit)
	fmt.Println("ENDS")

}

func sender(odd, even, quit chan<- int) {
	for x := 0; x <= 100; x++ {
		if x%2 == 0 {
			even <- x
		} else {
			odd <- x
		}
	}
	quit <- 0

}

func reciever(odd, even, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("from even channel", v)
		case v := <-odd:
			fmt.Println("from odd channel", v)
		case v := <-quit:
			fmt.Println("from quit  ", v)
			return //necessary so that select stops waiting for more
		}
	}

}
