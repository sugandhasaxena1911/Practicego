package main

import "fmt"

func main() {
	fmt.Println("Intro to channels ")
	ch := make(chan int)
	fmt.Printf("%T\n", ch)
	go func() {
		ch <- 43 // The code will not move forward till we recieve this value
	}()
	fmt.Println("recieve value ", <-ch)
}
