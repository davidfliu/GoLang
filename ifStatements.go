package main 

import "fmt"

func main() {

	var num int = 10

	if (num < 20) {

		fmt.Println("Your number is less than 20.")
	}

	if (num > 30) {

		fmt.Println("Your number is greater than 30.")

	} else {

		fmt.Println("Your number is less than 30.")
	}
	

	fmt.Printf("Your number is: %d", num)
}