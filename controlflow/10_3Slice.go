package main

import "fmt"

func main() {

	a := [10]int{1, 4, 5, 8, 9, 8} // default 0 forall elements
	fmt.Println(len(a), cap(a))
	fmt.Println(a)
	b := a[3:6] //cap of b = cap of a -no of elements sliced
	fmt.Println(b)
	fmt.Println(len(b), cap(b))
}
