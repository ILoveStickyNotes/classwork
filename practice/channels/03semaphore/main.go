package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//passes true to done clarifying that the anon func is done
		done <- true
	}()

	// If it wasn't in a go routine it wouldn't be able to complete
	go func() {
		//Waiting to be passed true value
		//After it's passed program can continue
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
