package main

import (
	"fmt"
)

func main() {
	c := factorial(6)
	for x := range c {
		fmt.Println(x)
	}

}

func factorial(x int) chan int {
	c := make(chan int)
	//Without the for loop being in a go routine it will deadlock
	go func() {
		total := 1
		for i := x; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()

	return c

}

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial
CHALLENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discussion area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others

package main

import (
	"fmt"
)

func main() {
	f := factorial(4)
	fmt.Println("Total:", f)
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}



*/
