package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)
	wg.Add(2)
	go func() {

		for x := 0; x <= 150; x++ {
			if x%2 == 0 {
				ch <- x

			}
		}

		wg.Done()
	}()
	go func() {

		for x := 0; x <= 150; x++ {
			if x%2 != 0 {
				ch <- x
			}
		}
		wg.Done()

	}()
	go func() {
		fmt.Println("waiting to close channel")
		wg.Wait()
		close(ch)
		fmt.Println(" channel closed")
	}()

	//go func() {   --> WHY code doest give output when his part is in go routie
	fmt.Println("about to print vales ")
	for y := range ch {
		fmt.Println(y)
	}
	fmt.Println("values printed ")
	//}()

	fmt.Println("EXITING")

}
