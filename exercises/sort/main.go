package main

import "sort"

type people []string

type peopleSorter struct {
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(peopleSorter)
}
