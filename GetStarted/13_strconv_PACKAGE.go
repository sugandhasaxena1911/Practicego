package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Atoi string to int) and Itoa (int to string).
	str := "34"
	num := 55
	i, err := strconv.Atoi(str)
	fmt.Println(i, err)
	fmt.Println(strconv.Itoa(num))

	//string to values
	b, err1 := strconv.ParseBool("true")
	j, _ := strconv.ParseInt("-1100", 2, 64)
	k, _ := strconv.ParseFloat("-15.3", 64)
	l, _ := strconv.ParseUint("765", 10, 64)
	fmt.Println(b, j, k, l, err1)
	//values to string
	s := strconv.FormatBool(true)
	s1 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s2 := strconv.FormatInt(999, 16)
	s3 := strconv.FormatUint(42, 2)
	fmt.Println(s, s1, s2, s3)

}
