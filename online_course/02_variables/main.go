package main

import "fmt"

func main() {

	// Print
	fmt.Println("Hello World!")

	// Variable
	var name string = "Tim"
	fmt.Println(name)

	var age int = 30
	fmt.Println(name, age)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)

	const constant string = "Detta kan ej ändras"

	newName := "Eric"
	fmt.Println(newName, age)
	fmt.Printf("%T\n", newName)

	size := 1.3
	fmt.Printf("%T\n", size)

	name2, email2 := "Tim", "tim@gmail.com"
	fmt.Println(name2, size, email2)
	fmt.Printf("%T\n", email2)

}
