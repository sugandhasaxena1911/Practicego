package main

import (
	"errors"
	"fmt"
	"github.com/sugandhasaxena1911/Practicego/GetStarted/MyError"
)

func main() {

	// create error using New() function
	myError1 := MyError.New("WRONG MESSAGE")
	fmt.Println(myError1)
	fmt.Printf("%v %T\n", myError1, myError1)
	fmt.Println(myError1.Error())

	myError2 := errors.New("WRONG MESSAGE  NEW")

	fmt.Println(myError2)
	fmt.Printf("%v %T\n", myError2, myError2)
	fmt.Println(myError2.Error())

}
