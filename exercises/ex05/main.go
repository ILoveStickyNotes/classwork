package main

import "fmt"

/*
Find the sum of all multiples of 5 below 1000
*/

func main() {
	fmt.Println(multiples(5))

}

func multiples(x int) int {

	var sum int
	for i := 1; i < 1000000; i++ {
		if 5*i < 1000 {
			sum += 5 * i //Adds the answer to 5*i to the sum count

		}
	}
	return sum
}
