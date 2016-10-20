package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()
	// Calculating the factorial gen into 2 processes
	//Fan out. Distributing the work across a number of goroutines/channels
	f1 := factorial(in)
	f2 := factorial(in)

	var count int
	//Fan in. Taking multiple channels into one
	for n := range merge(f1, f2) {
		count++
		fmt.Println(count, ": ", n)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100000; i++ {

			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
   CHALLENGE #1:
   -- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
   -- Use the "fan out / fan in" pattern
   CHALLENGE #2:
   WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
   -- While running the factorial computations, try to find how much of your resources are being used.
   -- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/
