package main

import "fmt"

func main() {

	fmt.Print(takesint(4))

}

func takesint(x int) (int, bool) {
	var truefalse bool
	if (x % 2) == 0 {
		truefalse = true
	} else {
		truefalse = false
	}

	return x / 2, truefalse

}
