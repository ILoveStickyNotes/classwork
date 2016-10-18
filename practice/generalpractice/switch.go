package mainfafa

import "fmt"

func main() {

	switch "bob" {
	case "Hello":
		fmt.Println("Hello")
	case "Goodbye":
		fmt.Println("Goodbye")
		fallthrough //Keeps it going
	case "How are you":
		fmt.Println("How are you")
	default:
		fmt.Println("Default")
	}

	Friendname := "al"
	// When nothing after switch everything has to equate to boolean
	switch {
	case len(Friendname) == 2:
		fmt.Println("2 letters")
	default:
		fmt.Println("Couldn't find anything")
	}

}
