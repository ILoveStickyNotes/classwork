package main

import (
	"errors"
	"log"
)

//Idiomatic way to do errors
var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
		//return 0, errors.New("norgate math: square root of negative number")
	}
	// implementation
	return 42, nil
}
