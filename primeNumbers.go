/*This program finds prime numbers within a certain range*/

package main 

import "fmt"

func main() {

	// Local variables
	var i, j int

	//Lower bounds of range
	var lower int = 2

	//Upper bounds of range
	var upper int = 50

	for i=lower; i < upper; i++ {
		for j=lower; j <= (i/j); j++ {
			if (i%j==0) {
				break; //factor found, so not prime
			}
		}
		if (j > (i/j)) {
			fmt.Printf("%d is prime\n", i)
		}
	} 
}