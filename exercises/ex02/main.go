package main

import "fmt"

func main() {
	fmt.Print(greatestnum(9.10, 99.1, 32.0, 9.11, 32000))
}

func greatestnum(x ...float64) float64 {
	highestnum := 0.0
	for _, y := range x {
		if y > highestnum {
			highestnum = y
		}
	}
	return highestnum
}
