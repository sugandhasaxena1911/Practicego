package main

import "fmt"

func main() {
	fmt.Println("Intro to channels ")
	ch := make(chan int)
	//fmt.Println(<-ch) //gives error since nothing to recieve
	ch <- 43 // The code will move forward till we recieve this valued
	fmt.Println("we put value to channel and noone to recieve", <-ch)
}
