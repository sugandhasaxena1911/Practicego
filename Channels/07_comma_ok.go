package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go sender1(odd, even, quit)
	reciever1(odd, even, quit)
	fmt.Println("ENDS")

}

func sender1(odd, even, quit chan<- int) {
	for x := 0; x <= 10; x++ {
		if x%2 == 0 {
			even <- x
		} else {
			odd <- x
		}
	}
	close(quit)
}

func reciever1(odd, even, quit <-chan int) {
	for {
		//A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
		select {
		case v := <-even:
			fmt.Println("from even channel", v)
		case v := <-odd:
			fmt.Println("from odd channel", v)
		case v, ok := <-quit:
			//Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
			//v, ok := <-ch
			//ok is false if there are no more values to receive and the channel is closed.
			// PS : comment return and then see infinte loop
			//REASON : ????
			if !ok {
				fmt.Println("I am false ", v, ok)
			} else {
				fmt.Println("I am True ", v, ok)
			}
			return

		}
	}

}
