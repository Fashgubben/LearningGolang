package main

import (
	"fmt"
)

func main() {

	// Arrays

	// Declare
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])

	// Arrays Declare and assign
	colorArr := [2]string{"Blue", "Green"}

	fmt.Println(colorArr)
	fmt.Println(colorArr[0])
	fmt.Println(colorArr[1])

	// Slices
	colorSlice := []string{"Green", "White", "Orange"}
	fmt.Println(colorSlice)
	fmt.Println(len(colorSlice))
}
