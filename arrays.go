package main 

import "fmt"

func main() {

	count := 5
	const size = 5
	var n [size] int
	var i, j int

	for i = 0; i < count; i++ {
		n[i] = i + 100
	}

	for j = 0; j < count; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}