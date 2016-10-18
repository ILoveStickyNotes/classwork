package main

import "fmt"

func main() {

	if (false && false) || (true && false) || (false && true) {
		fmt.Print("it worked") // returns true
	} else {
		fmt.Print("It didn't") // Returns false
	}

}
