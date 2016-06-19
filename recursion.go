// Simple program to calculate factorial with recursion in go.

package main

import "fmt"

func Factorial(i int)(result int) {

	//Terminating case.
	if (i > 1) {

		result = i * Factorial(i - 1)
		
		return result
	}

	return 1

}

//Main function.
func main() {

	i := 5

	fmt.Printf("Factorial of %d is: %d\n", i, Factorial(i))
}