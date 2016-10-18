package mainqwews

import "fmt"

func main() {
	//data := []float64{1,2,3,4,5,6,7} This would work as well
	//n := average(data...) Passes it in one at a time
	n := average(1, 2, 3, 4, 5, 6)
	fmt.Println(n)
}

func average(floats ...float64) float64 {
	fmt.Println(floats)

	var total float64
	// Like looping through an array without having to do it manually
	for _, v := range floats {
		total += v
		fmt.Println("hello")
	}
	return total / float64(len(floats))
}
