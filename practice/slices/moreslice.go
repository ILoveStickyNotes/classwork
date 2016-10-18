package main

import "fmt"

func main() {
	slice := make([]int, 1)
	fmt.Println(slice[0])
	slice[0] = 7
	fmt.Println(slice[0])
	slice[0]++ // Add 1 to the slice at that position
	fmt.Println(slice[0])
}
