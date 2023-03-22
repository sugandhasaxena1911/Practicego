package main

import "fmt"

func main() {
	// basic way of doing
	var kb = 1024 // 2^10
	var mb = 1024 * kb
	var gb = 1024 * mb
	fmt.Printf("the kb %v \t %b\n", kb, kb)
	fmt.Printf("the mb %v \t %b\n", mb, mb)
	fmt.Printf("the gb %v \t %b ", gb, gb)
	// using iota and bitshift
	const (
		_   = iota
		kkb = 1 << (iota * 10)
		mmb = 1 << (iota * 10)
		ggb = 1 << (iota * 10)
	)
	fmt.Printf("the kb %v \t %b\n", kkb, kkb)
	fmt.Printf("the mb %v \t %b\n", mmb, mmb)
	fmt.Printf("the gb %v \t %b ", ggb, ggb)
}
