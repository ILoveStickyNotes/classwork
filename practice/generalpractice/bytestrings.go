package mainff

import "fmt"

func main() {
	for i := 128; i <= 5100; i++ {
		fmt.Println(string(i), " ", []byte(string(i)))
	}

}
