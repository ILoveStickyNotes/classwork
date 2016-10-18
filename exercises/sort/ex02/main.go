package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"B", "Q", "A", "F", "G"}
	n := []int{2, 54, 2, 5, 7, 345, 3, 45}
	fmt.Println(s, n)
	//sort.StringSlice(s).Sort()
	sort.Sort(sort.StringSlice(s))
	sort.Sort(sort.IntSlice(n))
	fmt.Println(s, n)
}
