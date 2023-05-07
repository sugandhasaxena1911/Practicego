package main

import (
	"errors"
	"fmt"
)

func main() {
	n, e := F1(-8)
	fmt.Println(n, e)
	n, e = F2(-8)
	fmt.Println(n, e)

	n, e = F3(-16)
	fmt.Println(n, e)

	n, e = F4(-32)
	fmt.Println(n, e)
}

func F1(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Not allowed -ve number")
	}
	return n, nil
}
func F2(n int) (int, error) {
	if n < 0 {
		return 0, errors.New(fmt.Sprintln("Not allowed -ve number", n))
	}
	return n, nil
}

func F3(n int) (int, error) {
	if n < 0 {
		return 0, errors.New(fmt.Sprintf("Not allowed -ve number %b", n))
	}
	return n, nil
}

func F4(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("Not allowed -ve number %b", n)
	}
	return n, nil
}
