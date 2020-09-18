package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func sumOfNums(num1, num2 int) int {
	return num1 + num2
}

func main() {

	fmt.Println(greeting("Tim"))
	fmt.Println(sumOfNums(15, 15))

}
