package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string //`json:"-"` Removes this from being exported
	Age         int    //`json:"wisdom score"` Changes Age to Wisdom Score
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	//Marshal returns a byteslice of p1 and converts it into readable json
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
