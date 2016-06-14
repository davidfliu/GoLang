package main

import "fmt"

func main() {

	var a int = 0

	for a < 30 {

		if a == 10 {

			fmt.Printf("The value of a skipped: %d\n", a)
			a++
			continue
		}

		fmt.Printf("The value of a : %d\n", a)
		a++
	}
}