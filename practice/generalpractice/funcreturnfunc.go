package mainz

import "fmt"

func wrapper() func() int {
	x := 1

	return func() int {
		x++
		return x
		}
	}
}

func main() {

	firstfunc := wrapper()
	fmt.Println(firstfunc())
	fmt.Println(firstfunc())

}
