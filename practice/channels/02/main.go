package main

import "fmt"

func main() {
	//Make channel
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			//Hand value at variable i to the channel
			c <- i //func stops until the value is passed
		}
		//Close(c) no longer can put values onto the channel but can still receive
		close(c)
	}()
	//Keep receiving values until it is done
	//and nothing else is on the channel. Program cannot end until everything is gone
	for n := range c {
		fmt.Println(n)
		//When received it goes back up and repeats

	}

}
