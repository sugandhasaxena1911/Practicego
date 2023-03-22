package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Sugandhagan"
	str2 := "SugandhaSaxena"
	str3 := "Sugandha is a very good girl "

	// return string
	fmt.Println("Clone returns ", strings.Clone(str1))

	//The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
	fmt.Println("Compare returns ", strings.Compare(str1, str2))

	fmt.Println("Contains returns ", strings.Contains(str1, "gan"))

	fmt.Println("Count returns ", strings.Count(str1, "ga"))

	s1, s2, bl := strings.Cut(str1, "ga")
	fmt.Println(s1, s2, bl)

	//string to slice using whitespace
	s3 := strings.Fields(str3) // returns slice of strings
	fmt.Println(s3)

	fmt.Println("The index is at ", strings.Index(str1, "dha"))
	fmt.Println("The index is at ", strings.Index(str1, "Xa"))
	fmt.Println("The index is at ", strings.LastIndex(str3, "a "))

	//slice to string with separator
	str5 := []string{"My", "name ", "Sugandha", "saxena"}
	str6 := strings.Join(str5, ",")
	fmt.Println(str6)

}
