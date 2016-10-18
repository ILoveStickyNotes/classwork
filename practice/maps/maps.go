package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	delete(m, "k2")
	fmt.Println("map:", m)

	_, ok := m["k2"] //returns true/false and 1/0
	fmt.Println("ok?:", ok)
	// map[typeofkey]typeofvalue{} assoc with key
	n := map[string]int{"food": 1, "bar": 2}
	fmt.Println("map:", n)
	//You can use this
	var myGreeting = map[string]string{}
	// Or this
	//myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Good Morning."
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)

	newmap := map[string]string{
		"bob": "b",
		"aoa": "a",
	}

	//If you want to add something to that
	newmap["ANOTHER"] = "WOW"
	fmt.Println(newmap)
	//If you want to delete
	delete(newmap, "bob")
	fmt.Println(newmap)

}
