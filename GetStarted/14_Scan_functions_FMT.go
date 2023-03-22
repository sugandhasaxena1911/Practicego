package main

import "fmt"

func main() {
	//Scan()
	//space separated inputs .new lines count as space
	//It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.
	//EOF is Ctrl Z + Enter
	var a, b, c int
	fmt.Println("Enter 3 numbers")
	i, err := fmt.Scan(&a, &b, &c)
	fmt.Println(i, err)
	fmt.Println(a, b, c)
	fmt.Println("Hello")

	//Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.
	// EOF is Ctrl Z + Enter   OR  Newline(Enter)
	//thefore ctrl Z from previous will give error here and inputs wont be scanned ***************
	fmt.Println("Enter 3 numbers")
	i, err = fmt.Scanln(&a, &b, &c)
	fmt.Println(i, err)
	fmt.Println(a, b, c)
	fmt.Println("Hello")

	//Scanf scans text read from standard input, storing successive space-separated values into successive arguments
	//as determined by the format. It returns the number of items successfully scanned.
	//If that is less than the number of arguments, err will report why.
	//Newlines in the input must match newlines in the format.
	//The one exception: the verb %c always scans the next rune in the input, even if it is a space (or tab etc.) or newline.
	//thefore ctrl Z from previous will give error here and inputs wont be scanned ***************
	fmt.Println("Enter 3 numbers")
	i, err = fmt.Scanf("%v %v %v", &a, &b, &c)
	fmt.Println(i, err)
	fmt.Println(a, b, c)
	fmt.Println("Hello")
}
