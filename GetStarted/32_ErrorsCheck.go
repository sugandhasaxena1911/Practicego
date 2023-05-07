package main

import "fmt"

func main() {
	st := stt{-2, "hello sugandha erorr ", 15.5}
	n, err := st.FError()
	fmt.Println(n, err)

}

func (abc stt) FError() (int, error) {
	if abc.a < 0 {
		return abc.a, stt{abc.a, abc.b, abc.c}

	}
	return 0, nil
}

type stt struct {
	a int
	b string
	c float64
}

func (abc stt) Error() string {
	return fmt.Sprintln("Appended string ", abc.a, abc.b, abc.c)
}
