package mainesdf

import "fmt"

func makeGreeter() func() func() func() string {
	return func() func() func() string {
		return func() func() string {
			return func() string {
				return "Hey"
			}
		}
	}
}

func main() {

	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Println(&greet)
	/*
	     You can only put a function in a function like this
	   	greeting := func() {
	   		fmt.Println("Hello")
	   	}

	   	greeting()
	   	fmt.Printf("%T\n", greeting)
	*/
}
