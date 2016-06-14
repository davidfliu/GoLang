//This program finds the max number between two numbers.

package main 

import "fmt"

func main() {

	//Declare and initialize variable all at once.
	num1, num2 := 50, 10

	fmt.Printf("The max number is: %d\n",max(num1, num2))
}

func max(num1, num2 int) int {

	var result int

	if num1 > num2 {
		result = num1;
	} else {
		result = num2;
	}

	return result
}