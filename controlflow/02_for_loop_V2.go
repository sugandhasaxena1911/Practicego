package main

import "fmt"

func main() {

	//method 1
	i := 1
	for i <= 5 {
		fmt.Println("Value", i)
		i++
	}
	//method 2
	j := 1
	for {
		if j > 5 {
			break
		}
		fmt.Println("Value", j)
		j++

	}

}
