package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		fmt.Println("Inside")
		for x := 0; x <= 150; x++ {
			if x%2 == 0 {
				ch <- x

			}
		}
	}()
	go func() {
		for {
			y := <-ch
			y++
			fmt.Println(y)
		}
	}()
	time.Sleep(time.Second) // : you need this to get output or waitgrp
}
