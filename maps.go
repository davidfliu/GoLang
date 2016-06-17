//Learning about the maps data type.
//Maps keys to values.

package main 

import "fmt"

func main() {

	//variable declaration with make function.
	var mapVar map[string]string
	mapVar = make(map[string]string)

	//Key-value pairs.
	mapVar["Apples"] = "Red"
	mapVar["Bananas"] = "Yellow"
	mapVar["Oranges"] = "Orange"

	//Delete key-value pair.
	delete(mapVar,"Oranges")

	//Print out all key-value pairs.
	for object := range mapVar {
		fmt.Println("Color of", object, "is", mapVar[object])
	} 

}

