package maind

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	// Referencing the pointer
	var b *int = &a //*int <--- type: pointer to an int

	fmt.Println(b)
	fmt.Println(*b) // Gets the value at the memory address (DEREFERENCE)

	*b = 42 // The value at this address change it to 42
	fmt.Println(a)
}
