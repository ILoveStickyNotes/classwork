package main

import (
	"fmt"
	"sync"
	"time"
)

//Concurrency doing many things but only one at a time
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	// go foo() and go bar() launches both funcitons simultaneously
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
