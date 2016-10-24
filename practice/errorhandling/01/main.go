package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	//idiomatic way to error check
	if err != nil {
		//fmt.Println("error happened", err) //Prints to standard out
		//log.Println("err happened", err) //Same thing except gives date
		log.Fatalln(err) // Does not allow program to continue. Exits program
		//panic(err) // Gives additional info and runs deferred functions before closing
	}
}
