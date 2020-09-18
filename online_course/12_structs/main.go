package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// HasBityhday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{firstName: "Tim", lastName: "O'Flaherty", city: "Stockholm", gender: "m", age: 30}
	person2 := Person{"Tom", "O'Riley", "Dublin", "m", 58}
	person3 := Person{"Sara", "Rosén Andersson", "Stockholm", "f", 33}

	fmt.Println(person1)
	fmt.Println(person1.firstName, person1.age)
	fmt.Println(person2)
	fmt.Println(person3)

	person1.hasBirthday()
	fmt.Println(person1.greet())

	person3.getMarried("O'Flaherty")
	fmt.Println(person3)

}
