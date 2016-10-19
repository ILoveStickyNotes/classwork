package main

import (
	"fmt"
)

func main() {
	c := multiplecomputes(numtochan(4, 6, 2, 3, 4, 5, 2, 3, 4, 34, 2, 5, 2, 34, 2, 99))

	for x := range c {
		fmt.Println(x)
	}
}

func numtochan(x ...int) chan int {
	c := make(chan int)
	//Without the c <- total being in a go routine it will deadlock
	go func() {
		for _, r := range x {
			c <- r
		}
		close(c)
	}()
	return c

}

func multiplecomputes(n chan int) chan int {
	factorials := make(chan int)
	go func() {

		for x := range n {
			factorials <- fact(x)
		}
		close(factorials)
	}()
	return factorials
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
CHALLENGE #1:
-- Change the code above to execute 100 factorial computations concurrently and in parallel.
-- Use the "pipeline" pattern to accomplish this
POST TO DISCUSSION:
-- What realizations did you have about working with concurrent code when building your solution?
-- eg, what insights did you have which helped you understand working with concurrency?
-- Post your answer, and read other answers, here: https://goo.gl/uJa99G
*/
