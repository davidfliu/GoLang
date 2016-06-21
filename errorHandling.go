package main 

import (

	"errors"
	"fmt"
	"math"
)

func Sqrt (value float64)(float64, error) {

	if (value < 0) {
		return 0, errors.New("Error: Negative number passed to Sqrt")
	}
	return math.Sqrt(value)
}

func main() {

	//Pass in false argument, to trigger error.
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//Pass in true argument, no error.
	result1, err1 := Sqrt(10)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}
}