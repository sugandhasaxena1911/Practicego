package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	n, e := fmt.Println("Hello")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("n", n)
	//scan
	var answer1, answer2 string
	fmt.Println("enter name ")
	n, e = fmt.Scan(&answer1)
	if e != nil {
		panic(e)
	}
	fmt.Println("enter second name ")

	n, e = fmt.Scan(&answer2)
	if e != nil {
		panic(e)
	}
	fmt.Println(answer1, answer2)

	fmt.Println("File ")
	f, e := os.Create("names.txt")
	if e != nil {
		fmt.Println("e")
		return
	}
	defer f.Close()

	r := strings.NewReader("My name is sugandha ")
	io.Copy(f, r)

	fmt.Println("open a file")
	f, e = os.Open("names.txt")
	if e != nil {
		fmt.Println("e")
		return
	}
	defer f.Close()
	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
