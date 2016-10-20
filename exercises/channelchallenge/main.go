package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var count int64
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incrementor("1")
	go incrementor("2")
	wg.Wait()
	fmt.Println("Final Counter:", count)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		mutex.Lock()
		count++
		mutex.Unlock()
		fmt.Println("Process: "+s+" printing:", i)
	}
	wg.Done()
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
