package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//Create logs.txt
	f, err := os.Create("logs.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f) // This line will set log output to the file

	//open a file that doesnt exist & write logs in file craeted above
	f2, err := os.Open("hello.txt")
	if err != nil {
		log.Println("the error occured while opening the Hello file ", err)
	}

	defer f2.Close()

}
