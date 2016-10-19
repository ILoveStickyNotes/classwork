package main

import "fmt"

func main() {
	//Assigning it to a variable makes it run
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

// chan int can send or receive channel types
func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
