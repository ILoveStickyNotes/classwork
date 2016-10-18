package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	// NewEncoder(whereyouarewritingitto).Encode(whatyou'rewriting)
	json.NewEncoder(os.Stdout).Encode(p1)
}
