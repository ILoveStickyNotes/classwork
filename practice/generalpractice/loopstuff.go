package mainzz

import "fmt"

func main() {

	i := 0
	//Continue makes it come back here
	for {
		i++
		if i%2 == 0 {
			continue //Goes back to the top
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}

}
