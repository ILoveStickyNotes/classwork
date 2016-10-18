package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	h, _ := http.Get("http://google.com")
	a, _ := ioutil.ReadAll(h.Body)
	h.Body.Close()
	fmt.Printf("%s", a)
	b := 

}
