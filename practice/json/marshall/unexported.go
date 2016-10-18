package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	//make sure to make the variables uppcase to export
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
