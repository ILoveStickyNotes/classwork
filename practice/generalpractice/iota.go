package mainb

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1
	MB = 1 << (iota * 10) // 2
	GB = 1 << (iota * 10) // 3
	A  = 1 << 3
)

func main() {

	fmt.Printf("%b\t %d\t", KB, KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%b\t %d\t", A, A) // In binary Adds a 0 in
}
