package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := `immyFAVOURITE`
	//func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(" error ", err)
	}
	fmt.Println("password", string(bs))
	fmt.Printf("%T%vtype\n", bs, bs)
	//compare
	err = bcrypt.CompareHashAndPassword(bs, []byte(s))
	if err != nil {
		fmt.Println("cannot login now, error occur ", err)
	}
	fmt.Println("Yo ! you are logged in successfully")
	s = `immyFAVOURITE1`
	err = bcrypt.CompareHashAndPassword(bs, []byte(s))
	if err != nil {
		fmt.Println("cannot login now")
	}
	fmt.Println("Yo ! you are logged in successfully")
}
