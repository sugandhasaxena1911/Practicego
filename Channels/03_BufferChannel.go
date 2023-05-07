package main

import "fmt"

func main() {
	ch := make(chan int, 1) // buffer channel with room for one value
	ch <- 3
	fmt.Println(<-ch)
}
