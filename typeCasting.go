//Small program to demonstrate type casting.

package main 

import "fmt"

func main() {

	var sum int = 10
	var count int = 15
	var mean float32

	//Typecasting from int to float32.
	mean = float32(sum)/float32(count)
	fmt.Printf("Value of mean is: %f\n", mean)
}