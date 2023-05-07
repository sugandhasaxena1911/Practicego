package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	type Employee struct {
		Firstname string
		Lastname  string
		Id        int
		Skills    []string
	}

	emp1 := Employee{"sugandha", "saxena", 1, []string{"Go", "database"}}
	fmt.Println(emp1)

	//converting thsi to json
	b, e := json.Marshal(emp1)
	if e != nil {
		log.Panicln("Cannot convert to json ", e)
	}
	fmt.Println(string(b))
	fmt.Println(b)
	os.Stdout.Write(b)

	// string to json
	st := "Sugandha"
	b, e = json.Marshal(st)
	fmt.Println("\nstring to json ", string(b))
	n := 500.5
	b, e = json.Marshal(n)
	fmt.Println(b)
	fmt.Println(string(b))

	/// slice of employee to json
	emps := []Employee{{"manu", "saxena", 2, []string{"Marketing", "Cartoons"}}, {"tina", "hello", 3, []string{"yelo", "yelo2"}}, emp1}
	b, e = json.Marshal(emps)
	if e != nil {
		log.Panicln(e)
	}
	fmt.Println("\nslice of employee to JSON ", string(b))
	os.Stdout.Write(b)
	fmt.Println("UNMARSHAL")
	//[]     : slice /JSON array
	//{} : struct/Json object
	//there below is : slice of structs
	s := `[{"Firstname":"sugandha","Lastname":"saxena"},{"Firstname":"manu", "Lastname":"saxena"}]`
	bs := []byte(s)
	type names struct {
		Firstname string `json:"firstname"`
		LastName  string `json:"lastname"`
	}
	familynames := []names{}
	e = json.Unmarshal(bs, &familynames)
	if e != nil {
		log.Println(e)
	}
	fmt.Println(familynames)
}
