package main

import "fmt"

func main() {
	for i := 32; i <= 133; i++ {

		fmt.Printf("the number %v has ascii %#U and hexdecimal %#x\n", i, i, i)
		//0-9   :  48 to 57
		//65-90 :  A-Z
		//97-122 : a-b
	}
}
