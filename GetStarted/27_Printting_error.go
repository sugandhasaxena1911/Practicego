package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//1. direct printting using fmt.pritln
	//2. using .log.Println   : this prints the timestamp as well
	f, err := os.Open("namez.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	fmt.Println("%T", f)

	//3. Log the logs into a file  , you may comment this part to print in standard library
	f1, e1 := os.Create("logs.txt")
	if e1 != nil {
		fmt.Println("err happened ", e1)
	}
	defer f1.Close()
	log.SetOutput(f1)

	//open a file again
	defer CloseMe2()
	f2, e2 := os.Open("namez2.txt")
	if e2 != nil {
		log.Println(e2)
	}
	defer f2.Close()

	//4. Panic
	defer CloseMe4()
	f4, e4 := os.Open("namez4.txt")
	if e4 != nil {
		log.Panicln(e4)
		//log.Println(e3)
	}
	defer f4.Close()
	fmt.Println("Am i Printed ??????????after panic ") // not printed and closeme3 () not called

	//5. log.fatal   : function deferred before call of fatal are nt called, therefore CloseMe3 , 2 are not called
	//println folowed by  os.exit,  with exit status code printed
	/*
		defer CloseMe3()
		f3, e3 := os.Open("namez3.txt")
		if e3 != nil {
			log.Fatalln(e3)
			//log.Println(e3)
		}
		defer f3.Close()
		fmt.Println("Am i Printed ??????????after fatal ") // not printed and closeme3 () not called
	*/

}

func CloseMe3() {
	fmt.Println("CloseME Please 3")
}

func CloseMe2() {
	fmt.Println("CloseME Please 2")
}
func CloseMe4() {
	fmt.Println("CloseME Please 4")
}
