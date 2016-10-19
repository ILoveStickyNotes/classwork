package main

import (
	"fmt"
	"sync"
)

//Why fan in/out?
//Two cores can do the processing.. Faster

func main() {
	in := gen(2, 3)

	// FAN OUT
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume the merged output from multiple channels.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func gen(nums ...int) chan int {
	fmt.Printf("TYPE OF NUMS %T\n", nums) // just FYI

	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	fmt.Printf("TYPE OF CS: %T\n", cs) // just FYI

	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		//Without putting ch chan int into the go routine params
		//it would cause a race condition because when it loops
		//over it tries to access C while the first go routine is running
		//c assigns its value to ch
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
			//you have to put (c) here
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
FAN OUT
Multiple functions reading from the same channel until that channel is closed
FAN IN
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.
PATTERN
there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.
source:
https://blog.golang.org/pipelines
*/

/*
CHALLENGE #1
When know HOW to do fan out / fan in, but do we know WHY?
Why would we want to do fan out / fan in?
*/
