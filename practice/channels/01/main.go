package main

import (
	"fmt"
	"time"
)

func main() {
	//Make channel
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			//Hand value at variable i to the channel
			c <- i //func stops until the value is passed
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) //Waiting to receive value off of the channel
			//When received it goes back up and repeats
		}
	}()

	time.Sleep(time.Second)
}
