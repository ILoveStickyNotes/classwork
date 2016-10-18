package mainq

import "fmt"

func zero(x *int) {
	fmt.Println(x)
	*x = 0 // Dereferences the value at this address to 0
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)

}
