package main

import "fmt"

func main() {

	s := "Hello, Hawaii"
	bs := []byte(s)
	fmt.Printf("value and type %v %T\n \n", bs, bs)

	// iterating a  string, is like iterating a slice of byte
	for i, v := range s {
		fmt.Println(i, v)
		fmt.Printf("%T %T\n", v, s[i])
		fmt.Printf("%#U \n", v)
	}

	///append to slice
	sl := make([]int, 5, 6)
	fmt.Println(sl, len(sl), cap(sl))
	sl = append(sl, 5)
	fmt.Println(sl, len(sl), cap(sl))
	sl = append(sl, 9)
	fmt.Println(sl, len(sl), cap(sl))

}
