package main

import (
	"fmt"
	"sort"
)

func main() {

	b := []int{1, 2, 6, 7, 8, 1, 90, 56, 8}
	fmt.Println("Is slice sorted ?? ", sort.IntsAreSorted(b))
	sort.Ints(b)
	fmt.Println("Sorted slice ", b)
	s := []string{"Hello", "Tina", "is ", "My", "name", "Name", "100", "20", "3", "4999"}
	sort.Strings(s)
	// they ares orted nby ascii value
	//0-9   :  48 to 57
	//65-90 :  A-Z
	//97-122 : a-b
	fmt.Println(s)
	fmt.Println("Is slice sorted ?? ", sort.StringsAreSorted(s))

}
