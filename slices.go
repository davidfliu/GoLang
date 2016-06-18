//In go, we use slices in place of arrays to offer more power and convenience.
//Slices have no specified length. 

package main 

import "fmt"

func main() {

	//Slice literal.
	letters := [] string {"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(letters)

	//Slice literal using make built in function.
	s := make([] int, 5, 5)
	fmt.Printf("Slice s: %d\n", s)
	fmt.Printf("Length of slice s: %d\n", len(s))
	fmt.Printf("Capactiy of slice s: %d\n", cap(s))

	//Slicing an existing array.
	slice1 := letters[2:]
	//slice2 := letters[ :4]
	//slice3 := letters[:]

	fmt.Printf("Value of slice 1: %f\n", slice1)
}