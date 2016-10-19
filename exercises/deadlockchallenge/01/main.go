package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//Must always be in a go routine when passing something in
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)

	//Before Solved
	/*
		c := make(chan int)
		c <- 1
		fmt.Println(<-c)
	*/
}
