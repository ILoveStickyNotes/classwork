package main

import "fmt"

func main() {

	mySlice := []string{"a", "b", "c", "d", "e"} // Can initialize like this
	// make([]string,2,10)  This gives better performance (lower cost)
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println("myString"[2])

	mySlice = append(mySlice, "f")

	otherslice := []string{"f", "g"}
	mySlice = append(mySlice, otherslice...) //Appending otherslice to myslice with ...
	fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice)
	//Deleting from a slice
	mySlice = append(mySlice[:4], mySlice[7:]...)
	fmt.Println("deleted e", mySlice)

	// Different ways to make a slice
	var a []string // = nil
	b := []string{}
	// b cannot be used like this   b[0] = "bob" but you can append
	// b = append(b,"bob")
	c := make([]string, 0, 0)
}
